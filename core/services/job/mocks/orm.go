// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	context "context"

	job "github.com/SeerLink/seerlink/core/services/job"
	mock "github.com/stretchr/testify/mock"

	pipeline "github.com/SeerLink/seerlink/core/services/pipeline"

	postgres "github.com/SeerLink/seerlink/core/services/postgres"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CheckForDeletedJobs provides a mock function with given fields: ctx
func (_m *ORM) CheckForDeletedJobs(ctx context.Context) ([]int32, error) {
	ret := _m.Called(ctx)

	var r0 []int32
	if rf, ok := ret.Get(0).(func(context.Context) []int32); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimUnclaimedJobs provides a mock function with given fields: ctx
func (_m *ORM) ClaimUnclaimedJobs(ctx context.Context) ([]job.SpecDB, error) {
	ret := _m.Called(ctx)

	var r0 []job.SpecDB
	if rf, ok := ret.Get(0).(func(context.Context) []job.SpecDB); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.SpecDB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *ORM) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateJob provides a mock function with given fields: ctx, jobSpec, taskDAG
func (_m *ORM) CreateJob(ctx context.Context, jobSpec *job.SpecDB, taskDAG pipeline.TaskDAG) error {
	ret := _m.Called(ctx, jobSpec, taskDAG)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *job.SpecDB, pipeline.TaskDAG) error); ok {
		r0 = rf(ctx, jobSpec, taskDAG)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJob provides a mock function with given fields: ctx, id
func (_m *ORM) DeleteJob(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindJob provides a mock function with given fields: id
func (_m *ORM) FindJob(id int32) (job.SpecDB, error) {
	ret := _m.Called(id)

	var r0 job.SpecDB
	if rf, ok := ret.Get(0).(func(int32) job.SpecDB); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(job.SpecDB)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JobsV2 provides a mock function with given fields:
func (_m *ORM) JobsV2() ([]job.SpecDB, error) {
	ret := _m.Called()

	var r0 []job.SpecDB
	if rf, ok := ret.Get(0).(func() []job.SpecDB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]job.SpecDB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListenForDeletedJobs provides a mock function with given fields:
func (_m *ORM) ListenForDeletedJobs() (postgres.Subscription, error) {
	ret := _m.Called()

	var r0 postgres.Subscription
	if rf, ok := ret.Get(0).(func() postgres.Subscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(postgres.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListenForNewJobs provides a mock function with given fields:
func (_m *ORM) ListenForNewJobs() (postgres.Subscription, error) {
	ret := _m.Called()

	var r0 postgres.Subscription
	if rf, ok := ret.Get(0).(func() postgres.Subscription); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(postgres.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PipelineRunsByJobID provides a mock function with given fields: jobID, offset, size
func (_m *ORM) PipelineRunsByJobID(jobID int32, offset int, size int) ([]pipeline.Run, int, error) {
	ret := _m.Called(jobID, offset, size)

	var r0 []pipeline.Run
	if rf, ok := ret.Get(0).(func(int32, int, int) []pipeline.Run); ok {
		r0 = rf(jobID, offset, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipeline.Run)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int32, int, int) int); ok {
		r1 = rf(jobID, offset, size)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int32, int, int) error); ok {
		r2 = rf(jobID, offset, size)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RecordError provides a mock function with given fields: ctx, jobID, description
func (_m *ORM) RecordError(ctx context.Context, jobID int32, description string) {
	_m.Called(ctx, jobID, description)
}

// UnclaimJob provides a mock function with given fields: ctx, id
func (_m *ORM) UnclaimJob(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
