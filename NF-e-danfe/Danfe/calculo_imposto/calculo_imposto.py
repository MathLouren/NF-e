# calculo_imposto.py
from reportlab.lib.units import mm

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_calculo_imposto(c, x0, y0, largura_total):
    """
    Desenha a seção de cálculo do imposto seguindo as novas normas da reforma tributária 2026
    """
    altura_total = 45 * mm
    linha_altura = 9 * mm
    
    # Posicionamento abaixo do destinatário
    y0 = y0 - 68 * mm - 28 * mm  # abaixo do destinatário
    
    # Título da seção
    c.setFont(FONTE_BOLD, 7)
    c.drawString(x0 + 0*mm, y0 + 1*mm, "CÁLCULO DO IMPOSTO")
    
    # --- Linha 1: BASE DE CÁLCULO IVA | VALOR IVA | BASE DE CÁLCULO IVA ST | VALOR IVA ST ---
    linha1_y = y0
    col_base_iva = largura_total * 0.25
    col_valor_iva = largura_total * 0.25
    col_base_st = largura_total * 0.25
    col_valor_st = largura_total * 0.25
    
    c.rect(x0, linha1_y - linha_altura, largura_total, linha_altura)
    c.line(x0 + col_base_iva, linha1_y, x0 + col_base_iva, linha1_y - linha_altura)
    c.line(x0 + col_base_iva + col_valor_iva, linha1_y, x0 + col_base_iva + col_valor_iva, linha1_y - linha_altura)
    c.line(x0 + col_base_iva + col_valor_iva + col_base_st, linha1_y, x0 + col_base_iva + col_valor_iva + col_base_st, linha1_y - linha_altura)
    
    # Cabeçalhos
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha1_y - 2.5*mm, "BASE DE CÁLCULO IVA")
    c.drawString(x0 + col_base_iva + 1.5*mm, linha1_y - 2.5*mm, "VALOR IVA")
    c.drawString(x0 + col_base_iva + col_valor_iva + 1.5*mm, linha1_y - 2.5*mm, "BASE DE CÁLCULO IVA ST")
    c.drawString(x0 + col_base_iva + col_valor_iva + col_base_st + 1.5*mm, linha1_y - 2.5*mm, "VALOR IVA ST")
    
    # Valores
    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x0 + 1.5*mm, linha1_y - 6*mm, "R$ 1.000,00")
    c.drawString(x0 + col_base_iva + 1.5*mm, linha1_y - 6*mm, "R$ 250,00")
    c.drawString(x0 + col_base_iva + col_valor_iva + 1.5*mm, linha1_y - 6*mm, "R$ 0,00")
    c.drawString(x0 + col_base_iva + col_valor_iva + col_base_st + 1.5*mm, linha1_y - 6*mm, "R$ 0,00")
    
    # --- Linha 2: VALOR TOTAL DOS PRODUTOS | VALOR DO FRETE | VALOR DO SEGURO | OUTRAS DESPESAS ---
    linha2_y = linha1_y - linha_altura
    col_produtos = largura_total * 0.25
    col_frete = largura_total * 0.25
    col_seguro = largura_total * 0.25
    col_despesas = largura_total * 0.25
    
    c.rect(x0, linha2_y - linha_altura, largura_total, linha_altura)
    c.line(x0 + col_produtos, linha2_y, x0 + col_produtos, linha2_y - linha_altura)
    c.line(x0 + col_produtos + col_frete, linha2_y, x0 + col_produtos + col_frete, linha2_y - linha_altura)
    c.line(x0 + col_produtos + col_frete + col_seguro, linha2_y, x0 + col_produtos + col_frete + col_seguro, linha2_y - linha_altura)
    
    # Cabeçalhos
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha2_y - 2.5*mm, "VALOR TOTAL DOS PRODUTOS")
    c.drawString(x0 + col_produtos + 1.5*mm, linha2_y - 2.5*mm, "VALOR DO FRETE")
    c.drawString(x0 + col_produtos + col_frete + 1.5*mm, linha2_y - 2.5*mm, "VALOR DO SEGURO")
    c.drawString(x0 + col_produtos + col_frete + col_seguro + 1.5*mm, linha2_y - 2.5*mm, "OUTRAS DESPESAS")
    
    # Valores
    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x0 + 1.5*mm, linha2_y - 6*mm, "R$ 1.000,00")
    c.drawString(x0 + col_produtos + 1.5*mm, linha2_y - 6*mm, "R$ 15,00")
    c.drawString(x0 + col_produtos + col_frete + 1.5*mm, linha2_y - 6*mm, "R$ 0,00")
    c.drawString(x0 + col_produtos + col_frete + col_seguro + 1.5*mm, linha2_y - 6*mm, "R$ 0,00")
    
    # --- Linha 3: VALOR TOTAL DO IPI | VALOR TOTAL DA NOTA | BASE DE CÁLCULO IBS | VALOR IBS ---
    linha3_y = linha2_y - linha_altura
    col_ipi = largura_total * 0.25
    col_total_nota = largura_total * 0.25
    col_base_ibs = largura_total * 0.25
    col_valor_ibs = largura_total * 0.25
    
    c.rect(x0, linha3_y - linha_altura, largura_total, linha_altura)
    c.line(x0 + col_ipi, linha3_y, x0 + col_ipi, linha3_y - linha_altura)
    c.line(x0 + col_ipi + col_total_nota, linha3_y, x0 + col_ipi + col_total_nota, linha3_y - linha_altura)
    c.line(x0 + col_ipi + col_total_nota + col_base_ibs, linha3_y, x0 + col_ipi + col_total_nota + col_base_ibs, linha3_y - linha_altura)
    
    # Cabeçalhos
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha3_y - 2.5*mm, "VALOR TOTAL DO IPI")
    c.drawString(x0 + col_ipi + 1.5*mm, linha3_y - 2.5*mm, "VALOR TOTAL DA NOTA")
    c.drawString(x0 + col_ipi + col_total_nota + 1.5*mm, linha3_y - 2.5*mm, "BASE DE CÁLCULO IBS")
    c.drawString(x0 + col_ipi + col_total_nota + col_base_ibs + 1.5*mm, linha3_y - 2.5*mm, "VALOR IBS")
    
    # Valores
    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x0 + 1.5*mm, linha3_y - 6*mm, "R$ 0,00")
    c.drawString(x0 + col_ipi + 1.5*mm, linha3_y - 6*mm, "R$ 1.015,00")
    c.drawString(x0 + col_ipi + col_total_nota + 1.5*mm, linha3_y - 6*mm, "R$ 1.000,00")
    c.drawString(x0 + col_ipi + col_total_nota + col_base_ibs + 1.5*mm, linha3_y - 6*mm, "R$ 50,00")
    
    # --- Linha 4: BASE DE CÁLCULO CBS | VALOR CBS | VALOR TOTAL DOS IMPOSTOS ---
    linha4_y = linha3_y - linha_altura
    col_base_cbs = largura_total * 0.33
    col_valor_cbs = largura_total * 0.33
    col_total_impostos = largura_total * 0.34
    
    c.rect(x0, linha4_y - linha_altura, largura_total, linha_altura)
    c.line(x0 + col_base_cbs, linha4_y, x0 + col_base_cbs, linha4_y - linha_altura)
    c.line(x0 + col_base_cbs + col_valor_cbs, linha4_y, x0 + col_base_cbs + col_valor_cbs, linha4_y - linha_altura)
    
    # Cabeçalhos
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha4_y - 2.5*mm, "BASE DE CÁLCULO CBS")
    c.drawString(x0 + col_base_cbs + 1.5*mm, linha4_y - 2.5*mm, "VALOR CBS")
    c.drawString(x0 + col_base_cbs + col_valor_cbs + 1.5*mm, linha4_y - 2.5*mm, "VALOR TOTAL DOS IMPOSTOS")
    
    # Valores
    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x0 + 1.5*mm, linha4_y - 6*mm, "R$ 1.000,00")
    c.drawString(x0 + col_base_cbs + 1.5*mm, linha4_y - 6*mm, "R$ 20,00")
    c.drawString(x0 + col_base_cbs + col_valor_cbs + 1.5*mm, linha4_y - 6*mm, "R$ 320,00")
    
    # --- Linha 5: INFORMAÇÕES ADICIONAIS ---
    linha5_y = linha4_y - linha_altura
    c.rect(x0, linha5_y - linha_altura, largura_total, linha_altura)
    
    # Cabeçalho
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha5_y - 2.5*mm, "INFORMAÇÕES ADICIONAIS")
    
    # Informações
    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x0 + 1.5*mm, linha5_y - 6*mm, "IVA: 25% - IBS: 5% - CBS: 2%")
    
    return y0 - altura_total
    