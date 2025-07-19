# emitente.py
from reportlab.lib.units import mm
from reportlab.lib import colors
from reportlab.graphics.barcode import code128

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_emitente(c, x0, y0, largura_total):
    y0 = 260 * mm
    altura_total = 36 * mm

    # Cálculo das larguras das colunas: 40% | 20% | 40%
    col1 = largura_total * 0.4
    col2 = largura_total * 0.2
    col3 = largura_total * 0.4

    # Caixa externa
    c.rect(x0, y0 - altura_total, largura_total, altura_total)

    # Divisões verticais
    x_col2 = x0 + col1
    x_col3 = x_col2 + col2
    c.line(x_col2, y0 - altura_total, x_col2, y0)
    c.line(x_col3, y0 - altura_total, x_col3, y0)

    # --- Coluna 1: Emitente ---
    c.setFont(FONTE_BOLD, 13)
    c.drawCentredString(x0 + col1 / 2, y0 - 6*mm, "KEJINGBR E-")
    c.drawCentredString(x0 + col1 / 2, y0 - 10*mm, "COMMERCE CO.")
    c.drawCentredString(x0 + col1 / 2, y0 - 14*mm, "LTDA")

    c.setFont(FONTE_TXT, 8)
    emit_info = "Avenida da Liberdade, 113, numero 67 - Liberdade,\nSao Paulo, SP - CEP: 01503000 \nFone: 00119653666719"
    lines = emit_info.split("\n")
    for i, line in enumerate(lines):
        c.drawString(x0 + 2*mm, y0 - (19 + i*4.2)*mm, line)

    # --- Coluna 2: DANFE + Nº, SÉRIE, Folha ---
    c.setFont(FONTE_BOLD, 12)
    c.drawCentredString(x_col2 + col2 / 2, y0 - 5*mm, "DANFE")

    c.setFont(FONTE_TXT, 7)
    c.drawCentredString(x_col2 + col2 / 2, y0 - 9.5*mm, "Documento Auxiliar da")
    c.drawCentredString(x_col2 + col2 / 2, y0 - 13*mm, "Nota Fiscal Eletrônica")

    # Entrada/Saída (alinhado verticalmente ao centro da caixa)
    c.setFont(FONTE_TXT, 8)
    linha_entrada = "0: Entrada"
    linha_saida = "1: Saída"

    # Calcula posição vertical da caixa e centraliza as duas linhas
    caixa_x = x_col2 + 20*mm
    caixa_y = y0 - 23.5*mm
    caixa_w = 6.5*mm
    caixa_h = 6.5*mm

    # Vertical center base para os textos
    texto_base_y = caixa_y + caixa_h / 2 + 1*mm
    c.drawString(x_col2 + 2*mm, texto_base_y, linha_entrada)
    c.drawString(x_col2 + 2*mm, texto_base_y - 4*mm, linha_saida)

    # Caixa e número centralizado
    c.setFont(FONTE_BOLD, 10)
    c.rect(caixa_x, caixa_y, caixa_w, caixa_h)
    c.drawCentredString(caixa_x + caixa_w / 2, caixa_y + caixa_h / 2 - 3, "1")

    # Nº e SÉRIE
    c.setFont(FONTE_BOLD, 11)
    c.drawString(x_col2 + 2*mm, y0 - 28*mm, "Nº")
    c.drawString(x_col2 + 2*mm, y0 - 32*mm, "SÉRIE")

    c.setFont(FONTE_BOLD, 11)
    c.drawRightString(x_col2 + col2 - 2*mm, y0 - 28*mm, "000.115.594")
    c.drawRightString(x_col2 + col2 - 2*mm, y0 - 32*mm, "002")

    # Folha - centralizado na coluna 2
    c.setFont(FONTE_TXT, 7)
    folha_text = "Folha 1/2"
    folha_width = c.stringWidth(folha_text, FONTE_TXT, 7)
    c.drawString(x_col2 + (col2 - folha_width) / 2, y0 - 35*mm, folha_text)

    # --- Coluna 3: Código de barras, chave e texto ---
    h1 = altura_total * 0.4
    h2 = altura_total * 0.25
    h3 = altura_total * 0.35

    # Parte 1: Código de barras
    barcode_value = "12345678901234567890123456789012345678901234"
    barcode = code128.Code128(barcode_value, barHeight=12*mm, barWidth=0.4)
    barcode_x = x_col3 + (col3 - barcode.width) / 2
    barcode_y = y0 - h1 + (h1 - barcode.height) / 2
    barcode.drawOn(c, barcode_x, barcode_y)

    # Parte 2: Chave de Acesso
    y_chave_top = y0 - h1
    y_chave_bot = y_chave_top - h2

    # Linhas superior e inferior da chave de acesso
    c.line(x_col3, y_chave_top, x_col3 + col3, y_chave_top)
    c.line(x_col3, y_chave_bot, x_col3 + col3, y_chave_bot)

    c.setFont(FONTE_BOLD, 7)
    c.drawCentredString(x_col3 + col3 / 2, y_chave_top - 3*mm, "CHAVE DE ACESSO")
    c.setFont(FONTE_BOLD, 7.5)
    c.drawCentredString(x_col3 + col3 / 2, y_chave_top - 7*mm, "1234 5678 9012 3456 7890 1234 5678 9012 3456")

    # Parte 3: Texto final
    y_texto_top = y_chave_bot
    c.setFont(FONTE_BOLD, 7)
    c.drawCentredString(x_col3 + col3 / 2, y_texto_top - 2.5*mm, "Consulta de autenticidade no portal nacional da NF-e")
    c.drawCentredString(x_col3 + col3 / 2, y_texto_top - 6*mm, "www.nfe.fazenda.gov.br/portal")
    c.drawCentredString(x_col3 + col3 / 2, y_texto_top - 9.5*mm, "ou no site da Sefaz Autorizadora")
 
    # ---------- Linha 1: Natureza / Protocolo (60|40) ----------
    linha1_y, linha_h = y0 - altura_total - 8*mm, 8*mm
    col60 = largura_total*0.6
    c.rect(x0, linha1_y, largura_total, linha_h)
    c.line(x0+col60, linha1_y, x0+col60, linha1_y+linha_h)

    c.setFont(FONTE_TXT, 6.5)
    c.drawString(x0+1.5*mm, linha1_y+linha_h-2.6*mm, "NATUREZA DA OPERAÇÃO")
    c.setFont(FONTE_TXT, 8.5)
    c.drawString(x0+1.5*mm, linha1_y+1.5*mm, "Venda de mercadorias")

    c.setFont(FONTE_TXT, 6.5)
    c.drawString(x0+col60+2*mm, linha1_y+linha_h-2.6*mm, "PROTOCOLO DE AUTORIZAÇÃO DE USO")
    c.setFont(FONTE_TXT, 8.5)
    c.drawString(x0+col60+2*mm, linha1_y+1.5*mm, "233250053505317 09/02/2025 10:09:39")

    # ---------- Linha 2: IE | IE ST | CNPJ (33|34|33) ----------
    linha2_y = linha1_y - linha_h
    c.rect(x0, linha2_y, largura_total, linha_h)
    colA, colB = largura_total*0.33, largura_total*0.34
    c.line(x0+colA, linha2_y, x0+colA, linha2_y+linha_h)
    c.line(x0+colA+colB, linha2_y, x0+colA+colB, linha2_y+linha_h)

    c.setFont(FONTE_TXT, 6.5)
    c.drawString(x0+1.5*mm, linha2_y+linha_h-2.6*mm, "INSCRIÇÃO ESTADUAL")
    c.setFont(FONTE_TXT, 8.5)
    c.drawString(x0+1.5*mm, linha2_y+1.5*mm, "138573843110")

    c.setFont(FONTE_TXT, 6.5)
    c.drawString(x0+colA+2*mm, linha2_y+linha_h-2.6*mm, "INSC. ESTADUAL DO SUBST. TRIBUTÁRIO")
    c.setFont(FONTE_TXT, 8.5)
    c.drawString(x0+colA+2*mm, linha2_y+1.5*mm, "")

    c.setFont(FONTE_TXT, 6.5)
    c.drawString(x0+colA+colB+2*mm, linha2_y+linha_h-2.6*mm, "CNPJ")
    c.setFont(FONTE_TXT, 8.5)
    c.drawString(x0+colA+colB+2*mm, linha2_y+1.5*mm, "49.456.846/0001-20")