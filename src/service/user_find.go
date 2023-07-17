package service

import (
	"context"
	"gin-base/src/dao"
	querymodel "gin-base/src/model/query"
	responsemodel "gin-base/src/model/response"
	"sync"
)

// All ...
func (s userImpl) All(ctx context.Context, q querymodel.UserAll) (res responsemodel.UserAll) {

	var (
		wg = sync.WaitGroup{}
		d  = dao.User()
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		users := d.All(ctx, q)
		res.List = s.getListUser(ctx, users)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		res.Total = d.Count(ctx)
	}()

	wg.Wait()

	res.Limit = q.Limit

	return
}

// Detail ...
func (s userImpl) Detail(ctx context.Context, id string) (res responsemodel.UserDetail, err error) {
	var d = dao.User()
	user, err := d.FindByID(ctx, id)
	if err != nil {
		return
	}

	res = s.detail(ctx, user)
	return
}
