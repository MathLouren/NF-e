from reportlab.lib.units import mm
from reportlab.lib import colors

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_qr_code(c, x0, y0, largura_total, dados_qr=None):
    """
    Desenha a seção de QR Code e informações de pagamento no DANFE
    dados_qr: dicionário com dados do QR Code e pagamento
    """
    altura_total = 20 * mm
    linha_altura = 10 * mm
    
    # Título da seção
    c.setFont(FONTE_BOLD, 7)
    c.drawString(x0 + 0*mm, y0 + 1*mm, "QR CODE E INFORMAÇÕES DE PAGAMENTO")
    
    # --- Área do QR Code (simulado) ---
    qr_size = 15 * mm
    qr_x = x0 + largura_total - qr_size - 2*mm
    qr_y = y0 - qr_size - 2*mm
    
    # Desenhar área do QR Code (retângulo simulado)
    c.setStrokeColor(colors.gray)
    c.setLineWidth(0.5)
    c.rect(qr_x, qr_y, qr_size, qr_size)
    
    # Texto "QR Code" no centro
    c.setFont(FONTE_BOLD, 8)
    c.setFillColor(colors.gray)
    c.drawCentredString(qr_x + qr_size/2, qr_y + qr_size/2 - 3*mm, "QR")
    c.drawCentredString(qr_x + qr_size/2, qr_y + qr_size/2 - 6*mm, "CODE")
    c.setFillColor(colors.black)
    
    # --- Informações de Pagamento ---
    info_x = x0
    info_y = y0
    info_width = largura_total - qr_size - 4*mm
    
    # Linha 1: Forma de Pagamento
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(info_x + 1.5*mm, info_y - 2.5*mm, "FORMA DE PAGAMENTO:")
    
    c.setFont(FONTE_TXT, 6.3)
    if dados_qr and dados_qr.get("forma_pagamento"):
        c.drawString(info_x + 1.5*mm, info_y - 6*mm, dados_qr["forma_pagamento"])
    else:
        c.drawString(info_x + 1.5*mm, info_y - 6*mm, "BOLETO BANCÁRIO")
    
    # Linha 2: Condições de Pagamento
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(info_x + 1.5*mm, info_y - 10*mm, "CONDIÇÕES DE PAGAMENTO:")
    
    c.setFont(FONTE_TXT, 6.3)
    if dados_qr and dados_qr.get("condicoes_pagamento"):
        c.drawString(info_x + 1.5*mm, info_y - 13.5*mm, dados_qr["condicoes_pagamento"])
    else:
        c.drawString(info_x + 1.5*mm, info_y - 13.5*mm, "À VISTA - VENCIMENTO: 09/03/2025")
    
    # --- URL de Consulta ---
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(info_x + 1.5*mm, info_y - 17*mm, "CONSULTA:")
    
    c.setFont(FONTE_TXT, 6.3)
    c.drawString(info_x + 1.5*mm, info_y - 20.5*mm, "www.nfe.fazenda.gov.br/portal")
    
    return y0 - altura_total 