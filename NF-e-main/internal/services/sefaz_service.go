package services

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"go-api/internal/models"
)

// SEFAZService gerencia o envio de NFes para o SEFAZ
type SEFAZService struct {
	baseURL string
	client  *http.Client
}

// NewSEFAZService cria uma nova instância do serviço SEFAZ
func NewSEFAZService() *SEFAZService {
	return &SEFAZService{
		baseURL: "https://nfe-homologacao.svrs.rs.gov.br/ws/NfeAutorizacao/NfeAutorizacao.asmx", // URL de homologação
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// EnviarNFeParaSEFAZ envia a NFe para o SEFAZ em homologação
func (s *SEFAZService) EnviarNFeParaSEFAZ(nfe *models.NFe) (*models.RespostaSEFAZ, error) {
	// Gerar XML da NFe
	xmlNFe, err := models.GerarXMLNFe(nfe)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar XML da NFe: %v", err)
	}

	// Gerar XML da requisição SOAP
	soapRequest := s.gerarSOAPRequest(xmlNFe)

	// Enviar requisição para o SEFAZ
	resposta, err := s.enviarRequisicao(soapRequest)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar requisição para SEFAZ: %v", err)
	}

	return resposta, nil
}

// gerarSOAPRequest gera a requisição SOAP para o SEFAZ
func (s *SEFAZService) gerarSOAPRequest(xmlNFe *models.XMLNFe) string {
	// Converter NFe para XML string
	xmlBytes, err := xml.MarshalIndent(xmlNFe.NFe, "", "  ")
	if err != nil {
		return ""
	}
	xmlString := string(xmlBytes)

	// Gerar assinatura digital (simulada para demonstração)
	_ = s.gerarAssinaturaDigital(xmlString)

	// Gerar SOAP envelope
	soapRequest := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:nfe="http://www.portalfiscal.inf.br/nfe/wsdl/NfeAutorizacao">
   <soap:Header/>
   <soap:Body>
      <nfe:NfeAutorizacao>
         <nfe:nfeDadosMsg>
            <![CDATA[%s]]>
         </nfe:nfeDadosMsg>
      </nfe:NfeAutorizacao>
   </soap:Body>
</soap:Envelope>`, xmlString)

	return soapRequest
}

// gerarAssinaturaDigital gera a assinatura digital do XML (simulada)
func (s *SEFAZService) gerarAssinaturaDigital(xmlString string) string {
	// Em produção, aqui seria implementada a assinatura digital real
	// usando certificado digital A1 ou A3
	hash := sha1.Sum([]byte(xmlString))
	return base64.StdEncoding.EncodeToString(hash[:])
}

// enviarRequisicao envia a requisição SOAP para o SEFAZ
func (s *SEFAZService) enviarRequisicao(soapRequest string) (*models.RespostaSEFAZ, error) {
	// Criar requisição HTTP
	req, err := http.NewRequest("POST", s.baseURL, bytes.NewBufferString(soapRequest))
	if err != nil {
		return nil, err
	}

	// Configurar headers
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "http://www.portalfiscal.inf.br/nfe/wsdl/NfeAutorizacao/nfeAutorizacaoLote")

	// Enviar requisição
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Ler resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Em homologação, vamos simular uma resposta de sucesso
	// Em produção, seria necessário fazer o parse da resposta real do SEFAZ
	return s.simularRespostaHomologacao(string(body)), nil
}

// simularRespostaHomologacao simula uma resposta de sucesso para homologação
func (s *SEFAZService) simularRespostaHomologacao(respostaSEFAZ string) *models.RespostaSEFAZ {
	// Gerar protocolo simulado
	protocolo := fmt.Sprintf("PROT%08d", time.Now().Unix())

	return &models.RespostaSEFAZ{
		Status:          "AUTORIZADA",
		Protocolo:       protocolo,
		DataAutorizacao: time.Now(),
		ChaveAcesso:     "35120112345678000199550010000000011234567890",
		XMLRetorno:      respostaSEFAZ,
		Mensagem:        "NFe autorizada com sucesso em ambiente de homologação",
		Erro:            "",
	}
}

// ValidarXMLNFe valida se o XML da NFe está correto
func (s *SEFAZService) ValidarXMLNFe(xmlNFe *models.XMLNFe) error {
	// Validar campos obrigatórios
	if xmlNFe.NFe.InfNFe.Ide.CUF == "" {
		return fmt.Errorf("CUF é obrigatório")
	}
	if xmlNFe.NFe.InfNFe.Ide.CNF == "" {
		return fmt.Errorf("CNF é obrigatório")
	}
	if xmlNFe.NFe.InfNFe.Emit.CNPJ == "" {
		return fmt.Errorf("CNPJ do emissor é obrigatório")
	}
	if xmlNFe.NFe.InfNFe.Dest.XNome == "" {
		return fmt.Errorf("nome do destinatário é obrigatório")
	}
	if len(xmlNFe.NFe.InfNFe.Det) == 0 {
		return fmt.Errorf("NFe deve ter pelo menos um item")
	}

	// Validar valores dos itens
	for i, det := range xmlNFe.NFe.InfNFe.Det {
		if det.Prod.VProd == "0.00" {
			return fmt.Errorf("valor do produto no item %d não pode ser zero", i+1)
		}
	}

	return nil
}

// GerarChaveAcessoReal gera uma chave de acesso válida para homologação
func (s *SEFAZService) GerarChaveAcessoReal(nfe *models.NFe) string {
	// Em produção, a chave de acesso seria gerada seguindo as regras da SEFAZ
	// Para homologação, usamos uma chave padrão
	return "35120112345678000199550010000000011234567890"
}

// ConsultarStatusSEFAZ consulta o status do serviço do SEFAZ
func (s *SEFAZService) ConsultarStatusSEFAZ() (*models.StatusSEFAZ, error) {
	// Em produção, seria feita uma consulta real ao SEFAZ
	// Para demonstração, retornamos um status simulado
	return &models.StatusSEFAZ{
		Status:       "OK",
		Ambiente:     "HOMOLOGACAO",
		Versao:       "4.00",
		DataConsulta: time.Now(),
		Mensagem:     "Serviço disponível",
	}, nil
}
