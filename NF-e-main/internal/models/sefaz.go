package models

import "time"

// RespostaSEFAZ representa a resposta do SEFAZ
type RespostaSEFAZ struct {
	Status          string    `json:"status"`
	Protocolo       string    `json:"protocolo"`
	DataAutorizacao time.Time `json:"data_autorizacao"`
	ChaveAcesso     string    `json:"chave_acesso"`
	XMLRetorno      string    `json:"xml_retorno"`
	Mensagem        string    `json:"mensagem"`
	Erro            string    `json:"erro,omitempty"`
}

// StatusSEFAZ representa o status do serviço do SEFAZ
type StatusSEFAZ struct {
	Status       string    `json:"status"`
	Ambiente     string    `json:"ambiente"`
	Versao       string    `json:"versao"`
	DataConsulta time.Time `json:"data_consulta"`
	Mensagem     string    `json:"mensagem"`
}

// EnvioSEFAZRequest representa a requisição para envio ao SEFAZ
type EnvioSEFAZRequest struct {
	NFeID string `json:"nfe_id" binding:"required"`
}

// EnvioSEFAZResponse representa a resposta do envio ao SEFAZ
type EnvioSEFAZResponse struct {
	Status          string    `json:"status"`
	Protocolo       string    `json:"protocolo"`
	DataAutorizacao time.Time `json:"data_autorizacao"`
	ChaveAcesso     string    `json:"chave_acesso"`
	Mensagem        string    `json:"mensagem"`
	XMLGerado       string    `json:"xml_gerado"`
	XMLRetorno      string    `json:"xml_retorno"`
}
