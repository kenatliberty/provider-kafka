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

	//l := logging.Logger()

	logger := logging.Logger.Info("info-ng")

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
				log:         logger,
			},
			name: "Working1",
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
				// the type
				log: tt.fields.log,
			}
			fmt.Print(c)
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
