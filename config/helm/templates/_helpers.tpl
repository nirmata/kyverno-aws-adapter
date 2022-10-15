{{/*
Expand the name of the chart.
*/}}
{{- define "nirmata-aws-adapter.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "nirmata-aws-adapter.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "nirmata-aws-adapter.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "nirmata-aws-adapter.labels" -}}
helm.sh/chart: {{ include "nirmata-aws-adapter.chart" . }}
{{ include "nirmata-aws-adapter.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "nirmata-aws-adapter.selectorLabels" -}}
app.kubernetes.io/name: {{ include "nirmata-aws-adapter.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Generate the dockerconfigjson value
*/}}
{{- define "nirmata-aws-adapter.dockerconfigjson" -}}
{{- $user_pwd_hashed := printf "%s:%s" .Values.registryConfig.username .Values.registryConfig.password | b64enc }}
{{- printf "{\"auths\":{\"ghcr.io\":{\"auth\":\"%s\"}}}" $user_pwd_hashed | b64enc }}
{{- end }}
