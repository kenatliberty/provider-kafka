package topic

import (
	"context"
	"fmt"
	"github.com/crossplane-contrib/provider-kafka/apis/topic/v1alpha1"
	"github.com/crossplane-contrib/provider-kafka/internal/clients/kafka"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/twmb/franz-go/pkg/kadm"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	controllerruntime "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

var dataTesting = []byte(`{
		"brokers": [ "kafka-dev-0.kafka-dev-headless:9092"],
	"sasl": {
	"mechanism": "PLAIN",
	"username": "user",
	"password": "P0Jn0ZK9MA"
	}
	}`)

func TestObserve(t *testing.T) {

	newAc, _ := kafka.NewAdminClient(dataTesting)

	type fields struct {
		kafkaClient *kadm.Client
	}
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}

	cases := map[string]struct {
		name    string
		fields  fields
		args    args
		want    managed.ExternalObservation
		wantErr bool
	}{
		"ObserveT1": {
			fields: fields{
				kafkaClient: newAc,
			},
			name: "testTopic-1",
			args: args{
				ctx: context.Background(),
				// need to look at styra and compare how they did it

				mg: &v1alpha1.Topic{},
			},
			want: managed.ExternalObservation{
				ResourceExists:          true,
				ResourceUpToDate:        true,
				ResourceLateInitialized: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			c := &external{
				kafkaClient: tt.fields.kafkaClient,
			}
			got, err := c.Observe(tt.args.ctx, tt.args.mg)
			fmt.Print(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Observe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Observe() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetup(t *testing.T) {
	type args struct {
		mgr controllerruntime.Manager
		l   logging.Logger
		rl  workqueue.RateLimiter
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Setup(tt.args.mgr, tt.args.l, tt.args.rl); (err != nil) != tt.wantErr {
				t.Errorf("Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_connector_Connect(t *testing.T) {
	type fields struct {
		kube         client.Client
		usage        resource.Tracker
		log          logging.Logger
		newServiceFn func(creds []byte) (*kadm.Client, error)
	}
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    managed.ExternalClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &connector{
				kube:         tt.fields.kube,
				usage:        tt.fields.usage,
				log:          tt.fields.log,
				newServiceFn: tt.fields.newServiceFn,
			}
			got, err := c.Connect(tt.args.ctx, tt.args.mg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_external_Create(t *testing.T) {
	type fields struct {
		kafkaClient *kadm.Client
		log         logging.Logger
	}
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    managed.ExternalCreation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &external{
				kafkaClient: tt.fields.kafkaClient,
				log:         tt.fields.log,
			}
			got, err := c.Create(tt.args.ctx, tt.args.mg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_external_Delete(t *testing.T) {
	type fields struct {
		kafkaClient *kadm.Client
		log         logging.Logger
	}
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &external{
				kafkaClient: tt.fields.kafkaClient,
				log:         tt.fields.log,
			}
			if err := c.Delete(tt.args.ctx, tt.args.mg); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_external_Observe(t *testing.T) {
	type fields struct {
		kafkaClient *kadm.Client
		log         logging.Logger
	}
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    managed.ExternalObservation
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &external{
				kafkaClient: tt.fields.kafkaClient,
				log:         tt.fields.log,
			}
			got, err := c.Observe(tt.args.ctx, tt.args.mg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Observe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Observe() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_external_Update(t *testing.T) {
	type fields struct {
		kafkaClient *kadm.Client
		log         logging.Logger
	}
	type args struct {
		ctx context.Context
		mg  resource.Managed
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    managed.ExternalUpdate
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &external{
				kafkaClient: tt.fields.kafkaClient,
				log:         tt.fields.log,
			}
			got, err := c.Update(tt.args.ctx, tt.args.mg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
