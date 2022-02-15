package topic

import (
	"context"
	"github.com/crossplane-contrib/provider-kafka/internal/clients/kafka"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/twmb/franz-go/pkg/kadm"
	"reflect"
	"testing"
)

var dataTesting = []byte(`{
		"brokers": [ "kafka-dev-0.kafka-dev-headless:9092"],
	"sasl": {
	"mechanism": "PLAIN",
	"username": "user",
	"password": ""
	}
	}`)

func TestObserve(t *testing.T) {

	newAc, _ := kafka.NewAdminClient(dataTesting)

	type fields struct {
		kafkaClient *kadm.Client
		log         logging.Logger
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
				log:         nil,
			},
			name: "Working1",
			args: args{
				ctx: context.Background(),
				mg: &Topic{
					Name:              "testTopic-2",
					ReplicationFactor: 1,
					Partitions:        1,
					Config:            nil,
				},
			},
			want: want{
				cr: aef,
				result: managed.ExternalObservation{
					ResourceExists:          true,
					ResourceUpToDate:        false,
					ResourceLateInitialized: false,
				},
			},
		},
	}
	for _, tt := range cases {
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
