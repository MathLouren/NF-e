from reportlab.lib.units import mm

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_faturas(c, x0, y0, largura_total, dados_faturas=None):
    """
    Desenha a seção de FATURAS/DUPLICATAS no DANFE
    dados_faturas: lista de dicionários com os dados das faturas (pode ser None para exemplo)
    """
    altura_total = 15 * mm
    linha_altura = 7 * mm
    
    # Título da seção
    c.setFont(FONTE_BOLD, 7)
    c.drawString(x0 + 0*mm, y0 + 1*mm, "FATURAS")
    
    # --- Cabeçalho da tabela ---
    linha_cabecalho_y = y0
    
    # Definição das colunas
    col_numero = largura_total * 0.20
    col_vencimento = largura_total * 0.25
    col_valor = largura_total * 0.25
    col_forma_pagamento = largura_total * 0.30
    
    # Posições X das colunas
    x_numero = x0
    x_vencimento = x_numero + col_numero
    x_valor = x_vencimento + col_vencimento
    x_forma_pagamento = x_valor + col_valor
    
    # Desenhar cabeçalho
    c.rect(x0, linha_cabecalho_y - linha_altura, largura_total, linha_altura)
    
    # Linhas verticais do cabeçalho
    c.line(x_vencimento, linha_cabecalho_y, x_vencimento, linha_cabecalho_y - linha_altura)
    c.line(x_valor, linha_cabecalho_y, x_valor, linha_cabecalho_y - linha_altura)
    c.line(x_forma_pagamento, linha_cabecalho_y, x_forma_pagamento, linha_cabecalho_y - linha_altura)
    
    # Textos do cabeçalho
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x_numero + 1.5*mm, linha_cabecalho_y - 2.5*mm, "NÚMERO")
    c.drawString(x_vencimento + 1.5*mm, linha_cabecalho_y - 2.5*mm, "VENCIMENTO")
    c.drawString(x_valor + 1.5*mm, linha_cabecalho_y - 2.5*mm, "VALOR")
    c.drawString(x_forma_pagamento + 1.5*mm, linha_cabecalho_y - 2.5*mm, "FORMA DE PAGAMENTO")
    
    # --- Linha da Fatura 1 ---
    linha1_y = linha_cabecalho_y - linha_altura
    c.rect(x0, linha1_y - linha_altura, largura_total, linha_altura)
    
    # Linhas verticais
    c.line(x_vencimento, linha1_y, x_vencimento, linha1_y - linha_altura)
    c.line(x_valor, linha1_y, x_valor, linha1_y - linha_altura)
    c.line(x_forma_pagamento, linha1_y, x_forma_pagamento, linha1_y - linha_altura)
    
    # Dados da fatura 1
    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x_numero + 1.5*mm, linha1_y - 2.5*mm, "001")
    c.drawString(x_vencimento + 1.5*mm, linha1_y - 2.5*mm, "09/03/2025")
    c.drawString(x_valor + 1.5*mm, linha1_y - 2.5*mm, "R$ 1.300,00")
    c.drawString(x_forma_pagamento + 1.5*mm, linha1_y - 2.5*mm, "BOLETO BANCÁRIO")
    
    return y0 - altura_total 