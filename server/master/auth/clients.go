package auth

import (
	"fmt"
	"log"
	"math/rand"
	"net/rpc"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Clients struct {
	sync.RWMutex
	clients       map[string]*rpc.Client
	addresses     []string
	initConnsFlag int32
}

func NewClients(addresses []string) *Clients {
	cs := &Clients{}
	cs.addresses = addresses
	cs.clients = make(map[string]*rpc.Client)
	cs.initConnsFlag = 0
	return cs
}

func (cs *Clients) hasInit() bool {
	return atomic.LoadInt32(&cs.initConnsFlag) == 1
}

func (cs *Clients) setInitConnsFlag(val int32) {
	atomic.StoreInt32(&cs.initConnsFlag, val)
}

func (this *Clients) initConns() error {
	addresses := this.GetAddresses()
	count := len(addresses)

	if count == 0 {
		return fmt.Errorf("backend is blank")
	}

	var ee error
	for i := 0; i < count; i++ {
		endpoint := addresses[i]
		client, err := NewClient("tcp", endpoint, time.Second)
		if err != nil {
			log.Println("[F] cannot connect to", endpoint)
			ee = err
		}
		this.clients[endpoint] = client
	}

	return ee
}

// 这里的做法很简单，addresses列表一旦初始化设置好，就不变了，所以获取的时候也无需加锁
// 以后可以考虑做一个健康检查自动摘掉坏的实例，那个时候就要加锁了，复杂性增加，以后再说
func (this *Clients) GetAddresses() []string {
	return this.addresses
}

func (this *Clients) SetAddresses(addresses []string) {
	this.addresses = addresses
}

func (this *Clients) SetClients(clients map[string]*rpc.Client) {
	this.Lock()
	this.clients = clients
	this.Unlock()
}

func (this *Clients) PutClient(addr string, client *rpc.Client) {
	this.Lock()
	c, has := this.clients[addr]
	if has && c != nil {
		c.Close()
	}

	this.clients[addr] = client
	this.Unlock()
}

func (this *Clients) GetClient(addr string) (*rpc.Client, bool) {
	this.RLock()
	c, has := this.clients[addr]
	this.RUnlock()
	return c, has
}

func (this *Clients) Call(method string, args, reply interface{}) error {
	if !this.hasInit() {
		if err := this.initConns(); err != nil {
			return err
		}
		this.setInitConnsFlag(1)
	}

	addrs := this.GetAddresses()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, i := range r.Perm(len(addrs)) {
		addr := addrs[i]
		client, has := this.GetClient(addr)
		if !has {
			log.Println(addr, "has no client")
			continue
		}

		err := client.Call(method, args, reply)
		if err == nil {
			return nil
		}

		if err == rpc.ErrShutdown || strings.Contains(err.Error(), "connection refused") {
			// 后端可能重启了以至于原来持有的连接关闭，或者后端挂了
			// 可以尝试再次建立连接，搞定重启的情况
			client, err = NewClient("tcp", addr, time.Second)
			if err != nil {
				// 后端确实死翘翘了，继续尝试别的后端
				log.Println(addr, "is dead")
				continue
			} else {
				// 重新建立了与该实例的连接
				this.PutClient(addr, client)
				return client.Call(method, args, reply)
			}
		}

		// 刚开始后端没挂，但是仍然失败了，比如请求时间比较长，还没有结束，后端重启了，unexpected EOF
		// 不确定此时后端逻辑是否真的执行过了，防止后端逻辑不幂等，无法重试
		return err
	}

	return fmt.Errorf("all backends are dead")
}
