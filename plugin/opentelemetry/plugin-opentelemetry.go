package opentelemetry

import (
	ofctx "github.com/OpenFunction/functions-framework-go/context"
	"github.com/OpenFunction/functions-framework-go/plugin"
	"go.opentelemetry.io/otel/trace"
)

type PluginOpenTelemetry struct {
	tr trace.Tracer
}

func (p PluginOpenTelemetry) Name() string {
	//TODO implement me
	panic("implement me")
}

func (p PluginOpenTelemetry) Version() string {
	//TODO implement me
	panic("implement me")
}

func (p PluginOpenTelemetry) Init() plugin.Plugin {
	//TODO implement me
	panic("implement me")
}

func (p PluginOpenTelemetry) ExecPreHook(ofCtx ofctx.Context, plugins map[string]plugin.Plugin) error {
	if p.tr == nil {
		return nil
	}
	if ofCtx.SyncRequestMeta.Request != nil {
		request := ofCtx.SyncRequestMeta.Request

		ctx, _ := p.tr.Start(request.Context(), ofCtx.Name)

		ofCtx.SyncRequestMeta.Request = request.WithContext(ctx)
	}

	return nil
}

func (p PluginOpenTelemetry) ExecPostHook(ctx ofctx.Context, plugins map[string]plugin.Plugin) error {
	//TODO implement me
	panic("implement me")
}

func (p PluginOpenTelemetry) Get(fieldName string) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}
