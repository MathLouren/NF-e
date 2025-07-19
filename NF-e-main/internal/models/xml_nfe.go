package models

import (
	"encoding/xml"
	"fmt"
	"time"
)

// XMLNFe representa a estrutura XML da NFe conforme especificação da SEFAZ
type XMLNFe struct {
	XMLName   xml.Name     `xml:"nfeProc"`
	Versao    string       `xml:"versao,attr"`
	NFe       NFeXML       `xml:"NFe"`
	Protocolo ProtocoloXML `xml:"protNFe"`
}

// NFeXML representa a NFe no formato XML
type NFeXML struct {
	XMLName xml.Name  `xml:"NFe"`
	InfNFe  InfNFeXML `xml:"infNFe"`
}

// InfNFeXML representa as informações da NFe
type InfNFeXML struct {
	XMLName xml.Name `xml:"infNFe"`
	Id      string   `xml:"Id,attr"`
	Versao  string   `xml:"versao,attr"`
	Ide     IdeXML   `xml:"ide"`
	Emit    EmitXML  `xml:"emit"`
	Dest    DestXML  `xml:"dest"`
	Det     []DetXML `xml:"det"`
	Total   TotalXML `xml:"total"`
}

// IdeXML representa a identificação da NFe
type IdeXML struct {
	XMLName  xml.Name `xml:"ide"`
	CUF      string   `xml:"cUF"`
	CNF      string   `xml:"cNF"`
	NatOp    string   `xml:"natOp"`
	Mod      string   `xml:"mod"`
	Serie    string   `xml:"serie"`
	NNF      string   `xml:"nNF"`
	DhEmi    string   `xml:"dhEmi"`
	TpNF     string   `xml:"tpNF"`
	IdDest   string   `xml:"idDest"`
	CMunFG   string   `xml:"cMunFG"`
	TpImp    string   `xml:"tpImp"`
	TpEmis   string   `xml:"tpEmis"`
	CDV      string   `xml:"cDV"`
	TpAmb    string   `xml:"tpAmb"`
	FinNFe   string   `xml:"finNFe"`
	IndFinal string   `xml:"indFinal"`
	IndPres  string   `xml:"indPres"`
	ProcEmi  string   `xml:"procEmi"`
	VerProc  string   `xml:"verProc"`
}

// EmitXML representa o emissor da NFe
type EmitXML struct {
	XMLName   xml.Name     `xml:"emit"`
	CNPJ      string       `xml:"CNPJ"`
	XNome     string       `xml:"xNome"`
	XFant     string       `xml:"xFant,omitempty"`
	EnderEmit EnderEmitXML `xml:"enderEmit"`
	IE        string       `xml:"IE"`
	IM        string       `xml:"IM,omitempty"`
	CRT       string       `xml:"CRT"`
}

// EnderEmitXML representa o endereço do emissor
type EnderEmitXML struct {
	XMLName xml.Name `xml:"enderEmit"`
	XLgr    string   `xml:"xLgr"`
	Nro     string   `xml:"nro"`
	XCpl    string   `xml:"xCpl,omitempty"`
	XBairro string   `xml:"xBairro"`
	CMun    string   `xml:"cMun"`
	XMun    string   `xml:"xMun"`
	UF      string   `xml:"UF"`
	CEP     string   `xml:"CEP"`
	CPais   string   `xml:"cPais"`
	XPais   string   `xml:"xPais"`
}

// DestXML representa o destinatário da NFe
type DestXML struct {
	XMLName   xml.Name     `xml:"dest"`
	CNPJ      string       `xml:"CNPJ,omitempty"`
	CPF       string       `xml:"CPF,omitempty"`
	XNome     string       `xml:"xNome"`
	EnderDest EnderDestXML `xml:"enderDest"`
	IndIEDest string       `xml:"indIEDest"`
	IE        string       `xml:"IE,omitempty"`
}

// EnderDestXML representa o endereço do destinatário
type EnderDestXML struct {
	XMLName xml.Name `xml:"enderDest"`
	XLgr    string   `xml:"xLgr"`
	Nro     string   `xml:"nro"`
	XCpl    string   `xml:"xCpl,omitempty"`
	XBairro string   `xml:"xBairro"`
	CMun    string   `xml:"cMun"`
	XMun    string   `xml:"xMun"`
	UF      string   `xml:"UF"`
	CEP     string   `xml:"CEP"`
	CPais   string   `xml:"cPais"`
	XPais   string   `xml:"xPais"`
}

// DetXML representa um item da NFe
type DetXML struct {
	XMLName xml.Name   `xml:"det"`
	NItem   string     `xml:"nItem,attr"`
	Prod    ProdXML    `xml:"prod"`
	Imposto ImpostoXML `xml:"imposto"`
}

// ProdXML representa o produto/serviço
type ProdXML struct {
	XMLName  xml.Name `xml:"prod"`
	CProd    string   `xml:"cProd"`
	CEAN     string   `xml:"cEAN,omitempty"`
	XProd    string   `xml:"xProd"`
	NCM      string   `xml:"NCM"`
	CFOP     string   `xml:"CFOP"`
	UCom     string   `xml:"uCom"`
	QCom     string   `xml:"qCom"`
	VUnCom   string   `xml:"vUnCom"`
	VProd    string   `xml:"vProd"`
	CEANTrib string   `xml:"cEANTrib,omitempty"`
	UTrib    string   `xml:"uTrib"`
	QTrib    string   `xml:"qTrib"`
	VUnTrib  string   `xml:"vUnTrib"`
}

// ImpostoXML representa os impostos do item
type ImpostoXML struct {
	XMLName xml.Name  `xml:"imposto"`
	ICMS    ICMSXML   `xml:"ICMS"`
	PIS     PISXML    `xml:"PIS"`
	COFINS  COFINSXML `xml:"COFINS"`
}

// ICMSXML representa o ICMS
type ICMSXML struct {
	XMLName xml.Name  `xml:"ICMS"`
	ICMS00  ICMS00XML `xml:"ICMS00"`
}

// ICMS00XML representa o ICMS 00
type ICMS00XML struct {
	XMLName xml.Name `xml:"ICMS00"`
	Orig    string   `xml:"orig"`
	CST     string   `xml:"CST"`
	ModBC   string   `xml:"modBC"`
	VBC     string   `xml:"vBC"`
	PICMS   string   `xml:"pICMS"`
	VICMS   string   `xml:"vICMS"`
}

// PISXML representa o PIS
type PISXML struct {
	XMLName xml.Name   `xml:"PIS"`
	PISAliq PISAliqXML `xml:"PISAliq"`
}

// PISAliqXML representa o PIS por alíquota
type PISAliqXML struct {
	XMLName xml.Name `xml:"PISAliq"`
	CST     string   `xml:"CST"`
	VBC     string   `xml:"vBC"`
	PPIS    string   `xml:"pPIS"`
	VPIS    string   `xml:"vPIS"`
}

// COFINSXML representa o COFINS
type COFINSXML struct {
	XMLName    xml.Name      `xml:"COFINS"`
	COFINSAliq COFINSAliqXML `xml:"COFINSAliq"`
}

// COFINSAliqXML representa o COFINS por alíquota
type COFINSAliqXML struct {
	XMLName xml.Name `xml:"COFINSAliq"`
	CST     string   `xml:"CST"`
	VBC     string   `xml:"vBC"`
	PCOFINS string   `xml:"pCOFINS"`
	VCOFINS string   `xml:"vCOFINS"`
}

// TotalXML representa os totais da NFe
type TotalXML struct {
	XMLName xml.Name   `xml:"total"`
	ICMSTot ICMSTotXML `xml:"ICMSTot"`
}

// ICMSTotXML representa os totais do ICMS
type ICMSTotXML struct {
	XMLName xml.Name `xml:"ICMSTot"`
	VBC     string   `xml:"vBC"`
	VICMS   string   `xml:"vICMS"`
	VProd   string   `xml:"vProd"`
	VPIS    string   `xml:"vPIS"`
	VCOFINS string   `xml:"vCOFINS"`
}

// ProtocoloXML representa o protocolo de autorização
type ProtocoloXML struct {
	XMLName xml.Name   `xml:"protNFe"`
	Versao  string     `xml:"versao,attr"`
	InfProt InfProtXML `xml:"infProt"`
}

// InfProtXML representa as informações do protocolo
type InfProtXML struct {
	XMLName  xml.Name `xml:"infProt"`
	Id       string   `xml:"Id,attr"`
	TpAmb    string   `xml:"tpAmb"`
	VerAplic string   `xml:"verAplic"`
	ChNFe    string   `xml:"chNFe"`
	DhRecbto string   `xml:"dhRecbto"`
	NProt    string   `xml:"nProt"`
	DigVal   string   `xml:"digVal"`
	CStat    string   `xml:"cStat"`
	XMotivo  string   `xml:"xMotivo"`
}

// GerarXMLNFe gera o XML da NFe
func GerarXMLNFe(nfe *NFe) (*XMLNFe, error) {
	// Gerar chave de acesso (simplificado para demonstração)
	chaveAcesso := gerarChaveAcesso(nfe)

	// Formatar data/hora
	dataEmissao := nfe.DataEmissao.Format("2006-01-02T15:04:05-07:00")

	// Preparar itens
	var dets []DetXML
	for i, item := range nfe.Itens {
		det := DetXML{
			NItem: fmt.Sprintf("%d", i+1),
			Prod: ProdXML{
				CProd:   item.Codigo,
				XProd:   item.Descricao,
				NCM:     "00000000", // NCM padrão para demonstração
				CFOP:    "5102",     // CFOP padrão para venda
				UCom:    item.UnidadeMedida,
				QCom:    fmt.Sprintf("%.2f", item.Quantidade),
				VUnCom:  fmt.Sprintf("%.2f", item.ValorUnitario),
				VProd:   fmt.Sprintf("%.2f", item.ValorTotal),
				UTrib:   item.UnidadeMedida,
				QTrib:   fmt.Sprintf("%.2f", item.Quantidade),
				VUnTrib: fmt.Sprintf("%.2f", item.ValorUnitario),
			},
			Imposto: ImpostoXML{
				ICMS: ICMSXML{
					ICMS00: ICMS00XML{
						Orig:  "0",  // Origem nacional
						CST:   "00", // ICMS 00
						ModBC: "0",  // Valor da operação
						VBC:   fmt.Sprintf("%.2f", item.ValorTotal),
						PICMS: fmt.Sprintf("%.2f", item.AliquotaICMS),
						VICMS: fmt.Sprintf("%.2f", (item.ValorTotal*item.AliquotaICMS)/100),
					},
				},
				PIS: PISXML{
					PISAliq: PISAliqXML{
						CST:  "01", // Operação tributável
						VBC:  fmt.Sprintf("%.2f", item.ValorTotal),
						PPIS: fmt.Sprintf("%.2f", item.AliquotaPIS),
						VPIS: fmt.Sprintf("%.2f", (item.ValorTotal*item.AliquotaPIS)/100),
					},
				},
				COFINS: COFINSXML{
					COFINSAliq: COFINSAliqXML{
						CST:     "01", // Operação tributável
						VBC:     fmt.Sprintf("%.2f", item.ValorTotal),
						PCOFINS: fmt.Sprintf("%.2f", item.AliquotaCOFINS),
						VCOFINS: fmt.Sprintf("%.2f", (item.ValorTotal*item.AliquotaCOFINS)/100),
					},
				},
			},
		}
		dets = append(dets, det)
	}

	// Determinar CPF/CNPJ do destinatário
	var cpfDest, cnpjDest string
	if nfe.Destinatario.TipoPessoa == "F" {
		cpfDest = nfe.Destinatario.CPFCNPJ
	} else {
		cnpjDest = nfe.Destinatario.CPFCNPJ
	}

	xmlNFe := &XMLNFe{
		Versao: "4.00",
		NFe: NFeXML{
			InfNFe: InfNFeXML{
				Id:     chaveAcesso,
				Versao: "4.00",
				Ide: IdeXML{
					CUF:      "35", // SP
					CNF:      nfe.Numero,
					NatOp:    "Venda de mercadoria",
					Mod:      "55", // NFe
					Serie:    nfe.Serie,
					NNF:      nfe.Numero,
					DhEmi:    dataEmissao,
					TpNF:     "1",       // Saída
					IdDest:   "1",       // Interna
					CMunFG:   "3550308", // São Paulo
					TpImp:    "1",       // Retrato
					TpEmis:   "1",       // Normal
					CDV:      "1",       // Dígito verificador
					TpAmb:    "2",       // Homologação
					FinNFe:   "1",       // Normal
					IndFinal: "1",       // Consumidor final
					IndPres:  "1",       // Operação presencial
					ProcEmi:  "0",       // Aplicativo do contribuinte
					VerProc:  "1.0",
				},
				Emit: EmitXML{
					CNPJ:  nfe.Remetente.CPFCNPJ,
					XNome: nfe.Remetente.Nome,
					XFant: nfe.Remetente.NomeFantasia,
					EnderEmit: EnderEmitXML{
						XLgr:    nfe.Remetente.Endereco.Logradouro,
						Nro:     nfe.Remetente.Endereco.Numero,
						XCpl:    nfe.Remetente.Endereco.Complemento,
						XBairro: nfe.Remetente.Endereco.Bairro,
						CMun:    "3550308", // São Paulo
						XMun:    nfe.Remetente.Endereco.Cidade,
						UF:      nfe.Remetente.Endereco.Estado,
						CEP:     nfe.Remetente.Endereco.CEP,
						CPais:   "1058", // Brasil
						XPais:   nfe.Remetente.Endereco.Pais,
					},
					IE:  nfe.Remetente.InscricaoEstadual,
					IM:  nfe.Remetente.InscricaoMunicipal,
					CRT: nfe.Remetente.RegimeTributario,
				},
				Dest: DestXML{
					CNPJ:  cnpjDest,
					CPF:   cpfDest,
					XNome: nfe.Destinatario.Nome,
					EnderDest: EnderDestXML{
						XLgr:    nfe.Destinatario.Endereco.Logradouro,
						Nro:     nfe.Destinatario.Endereco.Numero,
						XCpl:    nfe.Destinatario.Endereco.Complemento,
						XBairro: nfe.Destinatario.Endereco.Bairro,
						CMun:    "3550308", // São Paulo
						XMun:    nfe.Destinatario.Endereco.Cidade,
						UF:      nfe.Destinatario.Endereco.Estado,
						CEP:     nfe.Destinatario.Endereco.CEP,
						CPais:   "1058", // Brasil
						XPais:   nfe.Destinatario.Endereco.Pais,
					},
					IndIEDest: "1", // Contribuinte ICMS
					IE:        nfe.Destinatario.InscricaoEstadual,
				},
				Det: dets,
				Total: TotalXML{
					ICMSTot: ICMSTotXML{
						VBC:     fmt.Sprintf("%.2f", nfe.ValorTotal),
						VICMS:   fmt.Sprintf("%.2f", nfe.ValorICMS),
						VProd:   fmt.Sprintf("%.2f", nfe.ValorTotal),
						VPIS:    fmt.Sprintf("%.2f", nfe.ValorPIS),
						VCOFINS: fmt.Sprintf("%.2f", nfe.ValorCOFINS),
					},
				},
			},
		},
	}

	return xmlNFe, nil
}

// gerarChaveAcesso gera uma chave de acesso para a NFe (simplificado)
func gerarChaveAcesso(nfe *NFe) string {
	// Formato: UF + AAMM + CNPJ + Mod + Série + Número + Tipo Emissão + Código Numérico + DV
	// Para demonstração, vamos gerar uma chave simplificada
	uf := "35" // SP
	aamm := time.Now().Format("0601")
	cnpj := nfe.Remetente.CPFCNPJ
	mod := "55" // NFe
	serie := nfe.Serie
	numero := nfe.Numero
	tpEmis := "1"        // Normal
	codNum := "00000001" // Código numérico
	dv := "1"            // Dígito verificador

	return uf + aamm + cnpj + mod + serie + numero + tpEmis + codNum + dv
}
