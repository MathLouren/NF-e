# recibo.py
from reportlab.lib.units import mm
from reportlab.lib import colors
from reportlab.platypus import Paragraph
from reportlab.lib.styles import getSampleStyleSheet

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

styles = getSampleStyleSheet()
estilo = styles["Normal"]
estilo.fontName = FONTE_TXT
estilo.fontSize = 7
estilo.leading = 8

texto_original = (
    "Recebemos de HOBBY GOODS BAZAR E ELETRONICOS 2023 LTDa, "
    "os produtos constantes da nota fiscal indicada ao lado: "
    "Data de emissão: 09/02/2025, Valor Total: R$37,90, "
    "Destinatário: Matheus Lourenco Rua das Flores, 832 - Neves - Sao Goncalo/RJ"
)

def cortar_para_duas_linhas(texto, largura_util):
    palavras = texto.split()
    for i in range(len(palavras), 0, -1):
        tentativa = " ".join(palavras[:i])
        par = Paragraph(tentativa, estilo)
        w, h = par.wrap(largura_util, 9999)
        num_linhas = int(round(h / estilo.leading))
        if num_linhas <= 2:
            return par
    return Paragraph("", estilo)

def desenhar_bloco_recibo(c, x0, y0, x_div, page_w, margem_h):
    ALTURA_BLOCO = 23 * mm
    ALTURA_CAMPOS_INFERIORES = 12 * mm
    largura_util = x_div - (x0 + 1*mm)

    par = cortar_para_duas_linhas(texto_original, largura_util)
    par.drawOn(c, x0 + 1*mm, y0 - 6.5*mm)

    y_caixas_top = y0 - 11*mm
    y_caixas_base = y_caixas_top - ALTURA_CAMPOS_INFERIORES
    larg_total_esq = x_div - x0
    larg_col1 = larg_total_esq * 0.35
    larg_col2 = larg_total_esq - larg_col1

    # Caixa 1
    c.line(x0, y_caixas_base, x0, y_caixas_base + ALTURA_CAMPOS_INFERIORES)
    c.line(x0, y_caixas_base + ALTURA_CAMPOS_INFERIORES, x0 + larg_col1, y_caixas_base + ALTURA_CAMPOS_INFERIORES)
    c.line(x0 + larg_col1, y_caixas_base, x0 + larg_col1, y_caixas_base + ALTURA_CAMPOS_INFERIORES)

    # Caixa 2
    c.line(x0 + larg_col1, y_caixas_base + ALTURA_CAMPOS_INFERIORES, x_div, y_caixas_base + ALTURA_CAMPOS_INFERIORES)
    c.line(x_div, y_caixas_base, x_div, y_caixas_base + ALTURA_CAMPOS_INFERIORES)

    c.setFont(FONTE_BOLD, 6.5)
    col1_text = "DATA DE RECEBIMENTO"
    col2_text = "IDENTIFICACAO E ASSINATURA DO RECEBEDOR"
    col1_width = c.stringWidth(col1_text, FONTE_BOLD, 6.5)
    col2_width = c.stringWidth(col2_text, FONTE_BOLD, 6.5)
    texto_y = y_caixas_base + ALTURA_CAMPOS_INFERIORES - 2.5*mm
    c.drawString(x0 + (larg_col1 - col1_width)/2, texto_y, col1_text)
    c.drawString(x0 + larg_col1 + (larg_col2 - col2_width)/2, texto_y, col2_text)

    return y_caixas_base
