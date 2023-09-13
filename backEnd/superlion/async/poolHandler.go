//@program: superlion
//@author: yanjl
//@create: 2023-09-12 10:13
package async

import "errors"

const (
	Running int64 = 0
	Stopped int64 = 1
)

// 将任务放到 channel 中供协程进行任务处理
func (p *Pool) Submit(t *Task) error {
	if p.status == Stopped {
		return errors.New("协程池已关闭，不能提交任务")
	}

	p.Lock()
	defer p.Unlock()
	if len(p.chTask) == int(p.capacity) {
		return errors.New("协程池已满，不能接受新任务")
	}

	p.chTask <- t
	return nil
}
