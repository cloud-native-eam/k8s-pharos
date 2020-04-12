

{{- define "alexandria.namespace" -}}
  {{- if .Values.namespaceOverride -}}
    {{- .Values.namespaceOverride -}}
  {{- end -}}
{{- end -}}