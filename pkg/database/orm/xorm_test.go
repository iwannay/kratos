package orm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type StatsCounter struct {
	Id              int64  `json:"id" xorm:"autoincr 'id'"`
	Topic           string `json:"topic"`
	Subject         string `json:"subject"`
	Object          string `json:"object"`
	CountType       string `json:"count_type"`
	TimeWindow      int    `json:"time_window"`
	TimeUnit        string `json:"time_unit"`
	Enabled         bool   `json:"enabled"`
	TTL             int    `json:"ttl" xorm:"ttl"`
	IsPersistence   bool   `json:"is_persistence"`
	QPS             int    `json:"qps" xorm:"qps"`
	AlarmValue      int64  `json:"alarm_value"`
	InterceptValue  int64  `json:"intercept_value"`
	NotifyType      string `json:"notify_type"`
	NotifyReceivers string `json:"notify_receivers"`
	EnabledNotify   bool   `json:"enabled_notify"`

	CreatedAt time.Time `xorm:"created" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at"`
	DeletedAt time.Time `xorm:"deleted" json:"-"`
}

func (*StatsCounter) TableName() string {
	return "s_stats_counter"
}

func TestClient_Switch(t *testing.T) {
	opt := &Options{
		Driver: "mysql",
		Dsn:    "ztm:SHdev#$%123@tcp(10.9.15.250:3306)/stats_counter?charset=utf8",
	}

	c := New(opt)
	for {

		count, err := c.Count(&StatsCounter{})
		t.Log(count, err)

		var counts []StatsCounter
		assert.Nil(t, c.Find(&counts), "")
		// t.Logf("\n%+v\n", counts)

		count, err = c.Switch(opt).Count(&StatsCounter{})
		t.Log(count, err)
		time.Sleep(1 * time.Second)
		time.Sleep(1 * time.Second)
	}
}
