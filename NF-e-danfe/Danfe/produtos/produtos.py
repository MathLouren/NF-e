# produtos.py
from reportlab.lib.units import mm

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_produtos(c, x0, y0, largura_total, dados_produtos=None, tipo="produto"):
    """
    Desenha a seção de DADOS DOS PRODUTOS/SERVIÇOS seguindo as novas normas da reforma tributária 2026
    dados_produtos: lista de dicionários com os dados dos produtos (pode ser None para exemplo)
    tipo: "produto" ou "servico" para adaptar os campos
    """
    altura_total = 35 * mm
    linha_altura = 7 * mm
    
    # Título da seção adaptado ao tipo
    c.setFont(FONTE_BOLD, 7)
    if tipo.lower() == "servico":
        titulo = "DADOS DOS SERVIÇOS - REFORMA TRIBUTÁRIA 2026"
    else:
        titulo = "DADOS DOS PRODUTOS/SERVIÇOS - REFORMA TRIBUTÁRIA 2026"
    c.drawString(x0 + 0*mm, y0 + 1*mm, titulo)
    
    # --- Cabeçalho da tabela ---
    linha_cabecalho_y = y0
    
    # Definição das colunas
    col_codigo = largura_total * 0.08
    col_descricao = largura_total * 0.25
    col_classificacao = largura_total * 0.08  # NCM para produtos, LC116 para serviços
    col_cfop = largura_total * 0.06
    col_unidade = largura_total * 0.06
    col_quantidade = largura_total * 0.08
    col_valor_unit = largura_total * 0.10
    col_valor_total = largura_total * 0.10
    col_base_iva = largura_total * 0.09
    col_valor_iva = largura_total * 0.10
    
    # Posições X das colunas
    x_codigo = x0
    x_descricao = x_codigo + col_codigo
    x_classificacao = x_descricao + col_descricao
    x_cfop = x_classificacao + col_classificacao
    x_unidade = x_cfop + col_cfop
    x_quantidade = x_unidade + col_unidade
    x_valor_unit = x_quantidade + col_quantidade
    x_valor_total = x_valor_unit + col_valor_unit
    x_base_iva = x_valor_total + col_valor_total
    x_valor_iva = x_base_iva + col_base_iva
    
    # Desenhar cabeçalho
    c.rect(x0, linha_cabecalho_y - linha_altura, largura_total, linha_altura)
    
    # Linhas verticais do cabeçalho
    c.line(x_descricao, linha_cabecalho_y, x_descricao, linha_cabecalho_y - linha_altura)
    c.line(x_classificacao, linha_cabecalho_y, x_classificacao, linha_cabecalho_y - linha_altura)
    c.line(x_cfop, linha_cabecalho_y, x_cfop, linha_cabecalho_y - linha_altura)
    c.line(x_unidade, linha_cabecalho_y, x_unidade, linha_cabecalho_y - linha_altura)
    c.line(x_quantidade, linha_cabecalho_y, x_quantidade, linha_cabecalho_y - linha_altura)
    c.line(x_valor_unit, linha_cabecalho_y, x_valor_unit, linha_cabecalho_y - linha_altura)
    c.line(x_valor_total, linha_cabecalho_y, x_valor_total, linha_cabecalho_y - linha_altura)
    c.line(x_base_iva, linha_cabecalho_y, x_base_iva, linha_cabecalho_y - linha_altura)
    c.line(x_valor_iva, linha_cabecalho_y, x_valor_iva, linha_cabecalho_y - linha_altura)
    
    # Textos do cabeçalho adaptados ao tipo
    c.setFont(FONTE_BOLD, 6)
    c.drawString(x_codigo + 1*mm, linha_cabecalho_y - 2.5*mm, "CÓDIGO")
    c.drawString(x_descricao + 1*mm, linha_cabecalho_y - 2.5*mm, "DESCRIÇÃO")
    
    if tipo.lower() == "servico":
        c.drawString(x_classificacao + 1*mm, linha_cabecalho_y - 2.5*mm, "LC116")
    else:
        c.drawString(x_classificacao + 1*mm, linha_cabecalho_y - 2.5*mm, "NCM")
    
    c.drawString(x_cfop + 1*mm, linha_cabecalho_y - 2.5*mm, "CFOP")
    c.drawString(x_unidade + 1*mm, linha_cabecalho_y - 2.5*mm, "UN")
    c.drawString(x_quantidade + 1*mm, linha_cabecalho_y - 2.5*mm, "QTD")
    c.drawString(x_valor_unit + 1*mm, linha_cabecalho_y - 2.5*mm, "VL.UNIT")
    c.drawString(x_valor_total + 1*mm, linha_cabecalho_y - 2.5*mm, "VL.TOTAL")
    c.drawString(x_base_iva + 1*mm, linha_cabecalho_y - 2.5*mm, "BASE IVA")
    c.drawString(x_valor_iva + 1*mm, linha_cabecalho_y - 2.5*mm, "VALOR IVA")
    
    # --- Linha do Item 1 ---
    linha1_y = linha_cabecalho_y - linha_altura
    c.rect(x0, linha1_y - linha_altura, largura_total, linha_altura)
    
    # Linhas verticais
    c.line(x_descricao, linha1_y, x_descricao, linha1_y - linha_altura)
    c.line(x_classificacao, linha1_y, x_classificacao, linha1_y - linha_altura)
    c.line(x_cfop, linha1_y, x_cfop, linha1_y - linha_altura)
    c.line(x_unidade, linha1_y, x_unidade, linha1_y - linha_altura)
    c.line(x_quantidade, linha1_y, x_quantidade, linha1_y - linha_altura)
    c.line(x_valor_unit, linha1_y, x_valor_unit, linha1_y - linha_altura)
    c.line(x_valor_total, linha1_y, x_valor_total, linha1_y - linha_altura)
    c.line(x_base_iva, linha1_y, x_base_iva, linha1_y - linha_altura)
    c.line(x_valor_iva, linha1_y, x_valor_iva, linha1_y - linha_altura)
    
    # Dados do item 1 adaptados ao tipo
    c.setFont(FONTE_TXT, 6)
    c.drawString(x_codigo + 1*mm, linha1_y - 2.5*mm, "001")
    
    if tipo.lower() == "servico":
        c.drawString(x_descricao + 1*mm, linha1_y - 2.5*mm, "SERVIÇO DE CONSULTORIA")
        c.drawString(x_classificacao + 1*mm, linha1_y - 2.5*mm, "0107")
        c.drawString(x_cfop + 1*mm, linha1_y - 2.5*mm, "5933")
        c.drawString(x_unidade + 1*mm, linha1_y - 2.5*mm, "HORA")
        c.drawString(x_quantidade + 1*mm, linha1_y - 2.5*mm, "10")
        c.drawString(x_valor_unit + 1*mm, linha1_y - 2.5*mm, "R$ 100,00")
        c.drawString(x_valor_total + 1*mm, linha1_y - 2.5*mm, "R$ 1.000,00")
    else:
        c.drawString(x_descricao + 1*mm, linha1_y - 2.5*mm, "PRODUTO EXEMPLO 1")
        c.drawString(x_classificacao + 1*mm, linha1_y - 2.5*mm, "12345678")
        c.drawString(x_cfop + 1*mm, linha1_y - 2.5*mm, "5102")
        c.drawString(x_unidade + 1*mm, linha1_y - 2.5*mm, "UN")
        c.drawString(x_quantidade + 1*mm, linha1_y - 2.5*mm, "2")
        c.drawString(x_valor_unit + 1*mm, linha1_y - 2.5*mm, "R$ 500,00")
        c.drawString(x_valor_total + 1*mm, linha1_y - 2.5*mm, "R$ 1.000,00")
    
    c.drawString(x_base_iva + 1*mm, linha1_y - 2.5*mm, "R$ 1.000,00")
    c.drawString(x_valor_iva + 1*mm, linha1_y - 2.5*mm, "R$ 250,00")
    
    # --- Linha do Item 2 ---
    linha2_y = linha1_y - linha_altura
    c.rect(x0, linha2_y - linha_altura, largura_total, linha_altura)
    
    # Linhas verticais
    c.line(x_descricao, linha2_y, x_descricao, linha2_y - linha_altura)
    c.line(x_classificacao, linha2_y, x_classificacao, linha2_y - linha_altura)
    c.line(x_cfop, linha2_y, x_cfop, linha2_y - linha_altura)
    c.line(x_unidade, linha2_y, x_unidade, linha2_y - linha_altura)
    c.line(x_quantidade, linha2_y, x_quantidade, linha2_y - linha_altura)
    c.line(x_valor_unit, linha2_y, x_valor_unit, linha2_y - linha_altura)
    c.line(x_valor_total, linha2_y, x_valor_total, linha2_y - linha_altura)
    c.line(x_base_iva, linha2_y, x_base_iva, linha2_y - linha_altura)
    c.line(x_valor_iva, linha2_y, x_valor_iva, linha2_y - linha_altura)
    
    # Dados do item 2 adaptados ao tipo
    c.setFont(FONTE_TXT, 6)
    c.drawString(x_codigo + 1*mm, linha2_y - 2.5*mm, "002")
    
    if tipo.lower() == "servico":
        c.drawString(x_descricao + 1*mm, linha2_y - 2.5*mm, "SERVIÇO DE MANUTENÇÃO")
        c.drawString(x_classificacao + 1*mm, linha2_y - 2.5*mm, "0301")
        c.drawString(x_cfop + 1*mm, linha2_y - 2.5*mm, "5933")
        c.drawString(x_unidade + 1*mm, linha2_y - 2.5*mm, "HORA")
        c.drawString(x_quantidade + 1*mm, linha2_y - 2.5*mm, "3")
        c.drawString(x_valor_unit + 1*mm, linha2_y - 2.5*mm, "R$ 100,00")
        c.drawString(x_valor_total + 1*mm, linha2_y - 2.5*mm, "R$ 300,00")
    else:
        c.drawString(x_descricao + 1*mm, linha2_y - 2.5*mm, "PRODUTO EXEMPLO 2")
        c.drawString(x_classificacao + 1*mm, linha2_y - 2.5*mm, "87654321")
        c.drawString(x_cfop + 1*mm, linha2_y - 2.5*mm, "5102")
        c.drawString(x_unidade + 1*mm, linha2_y - 2.5*mm, "UN")
        c.drawString(x_quantidade + 1*mm, linha2_y - 2.5*mm, "1")
        c.drawString(x_valor_unit + 1*mm, linha2_y - 2.5*mm, "R$ 300,00")
        c.drawString(x_valor_total + 1*mm, linha2_y - 2.5*mm, "R$ 300,00")
    
    c.drawString(x_base_iva + 1*mm, linha2_y - 2.5*mm, "R$ 300,00")
    c.drawString(x_valor_iva + 1*mm, linha2_y - 2.5*mm, "R$ 75,00")
    
    # --- Linha de Totais ---
    linha_totais_y = linha2_y - linha_altura
    c.rect(x0, linha_totais_y - linha_altura, largura_total, linha_altura)
    
    # Linhas verticais dos totais
    c.line(x_valor_total, linha_totais_y, x_valor_total, linha_totais_y - linha_altura)
    c.line(x_base_iva, linha_totais_y, x_base_iva, linha_totais_y - linha_altura)
    c.line(x_valor_iva, linha_totais_y, x_valor_iva, linha_totais_y - linha_altura)
    
    # Totais
    c.setFont(FONTE_BOLD, 6)
    c.drawString(x0 + 1*mm, linha_totais_y - 2.5*mm, "TOTAIS:")
    
    if tipo.lower() == "servico":
        c.drawString(x_valor_total + 1*mm, linha_totais_y - 2.5*mm, "R$ 1.300,00")
        c.drawString(x_base_iva + 1*mm, linha_totais_y - 2.5*mm, "R$ 1.300,00")
        c.drawString(x_valor_iva + 1*mm, linha_totais_y - 2.5*mm, "R$ 325,00")
    else:
        c.drawString(x_valor_total + 1*mm, linha_totais_y - 2.5*mm, "R$ 1.300,00")
        c.drawString(x_base_iva + 1*mm, linha_totais_y - 2.5*mm, "R$ 1.300,00")
        c.drawString(x_valor_iva + 1*mm, linha_totais_y - 2.5*mm, "R$ 325,00")
    
    # --- Linha de Informações Adicionais ---
    linha_info_y = linha_totais_y - linha_altura
    c.rect(x0, linha_info_y - linha_altura, largura_total, linha_altura)
    
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha_info_y - 2.5*mm, "INFORMAÇÕES ADICIONAIS")
    
    c.setFont(FONTE_TXT, 6.3)
    if tipo.lower() == "servico":
        c.drawString(x0 + 1.5*mm, linha_info_y - 6*mm, "IVA: 25% - IBS: 5% - CBS: 2% - SERVIÇOS - REFORMA TRIBUTÁRIA 2026")
    else:
        c.drawString(x0 + 1.5*mm, linha_info_y - 6*mm, "IVA: 25% - IBS: 5% - CBS: 2% - PRODUTOS - REFORMA TRIBUTÁRIA 2026")
    
    return y0 - altura_total 