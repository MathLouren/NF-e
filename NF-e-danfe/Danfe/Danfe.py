# main.py
from reportlab.pdfgen import canvas
from reportlab.lib.pagesizes import A4
from reportlab.lib.units import mm
from reportlab.lib import colors
from recibo.recibo import desenhar_bloco_recibo
from emitente.emitente import desenhar_emitente
from remetente.remetente import desenhar_destinatario
from calculo_imposto.calculo_imposto import desenhar_calculo_imposto
from transportador.transportador import desenhar_transportador
from produtos.produtos import desenhar_produtos
from faturas.faturas import desenhar_faturas
from dados_adicionais.dados_adicionais import desenhar_dados_adicionais
from qr_code.qr_code import desenhar_qr_code

PDF_PATH = "Danfe/danfe_templates/danfe.pdf"
PAGE_W, PAGE_H = A4
MARGEM_H = 8 * mm
COL_DIREITA = 35 * mm

c = canvas.Canvas(PDF_PATH, pagesize=A4)
c.setStrokeColor(colors.black)
c.setLineWidth(1)

# === PÁGINA 1 ===
x0 = MARGEM_H
y0 = PAGE_H - MARGEM_H
x_div = PAGE_W - MARGEM_H - COL_DIREITA

ALTURA_BLOCO = 23 * mm
y1 = y0 - ALTURA_BLOCO

# Moldura externa
c.rect(x0, y1, PAGE_W - 2 * MARGEM_H, ALTURA_BLOCO)
c.line(x_div, y1, x_div, y0)

# Bloco do recibo
y_caixas_base = desenhar_bloco_recibo(c, x0, y0, x_div, PAGE_W, MARGEM_H)

# Bloco NF-e
c.setFont("Helvetica-Bold", 13)

# Calcular a largura do texto e a posição para centralizar
nf_text = "NF-e"
nf_width = c.stringWidth(nf_text, "Helvetica-Bold", 13)

# Desenhar o texto centralizado
c.drawString(x_div + (COL_DIREITA - nf_width) / 2, y0 - 5*mm, nf_text)

c.setFont("Helvetica-Bold", 12)
c.drawString(x_div + 2*mm, y0 - 12*mm, "Nº")
c.setFont("Helvetica", 12)
c.drawRightString(PAGE_W - MARGEM_H - 2*mm, y0 - 12*mm, "000.115.594")

c.setFont("Helvetica-Bold", 12)
c.drawString(x_div + 2*mm, y0 - 16.5*mm, "SÉRIE")
c.setFont("Helvetica", 12)
c.drawRightString(PAGE_W - MARGEM_H - 2*mm, y0 - 16.5*mm, "002")

# Linha tracejada DANFE
y_tracinho = y_caixas_base - 3*mm
c.setLineWidth(1)
c.setDash(2.5, 2)
c.line(MARGEM_H, y_tracinho, PAGE_W - MARGEM_H, y_tracinho)
c.setDash()

# Posicionamento das seções - PÁGINA 1
y_emitente = 270*mm

# Desenhar seções na ordem correta - PÁGINA 1
desenhar_emitente(c, x0=8*mm, y0=y_emitente, largura_total=194*mm)
desenhar_destinatario(c, x0=8*mm, y0=y_emitente, largura_total=194*mm)
desenhar_calculo_imposto(c, x0=8*mm, y0=y_emitente, largura_total=194*mm)

# Transportador posicionado abaixo do cálculo do imposto
y_transportador = y_emitente - 68*mm - 28*mm - 45*mm - 7*mm
desenhar_transportador(c, x0=8*mm, y0=y_transportador, largura_total=194*mm)

# Produtos/Serviços posicionados abaixo do transportador
y_produtos = y_transportador - 21*mm - 10*mm

# Para produtos: tipo="produto" (padrão)
# Para serviços: tipo="servico"
desenhar_produtos(c, x0=8*mm, y0=y_produtos, largura_total=194*mm, tipo="produto")

# === PÁGINA 2 ===
c.showPage()
c.setStrokeColor(colors.black)
c.setLineWidth(1)

# Cabeçalho da página 2
c.setFont("Helvetica-Bold", 10)
c.drawString(x0 + 1*mm, PAGE_H - MARGEM_H - 5*mm, "DANFE - PÁGINA 2/2")
c.drawString(x0 + 1*mm, PAGE_H - MARGEM_H - 12*mm, "NF-e Nº 000.115.594 - SÉRIE 002")

# Faturas posicionadas no topo da página 2
y_faturas_p2 = PAGE_H - MARGEM_H - 25*mm
desenhar_faturas(c, x0=8*mm, y0=y_faturas_p2, largura_total=194*mm)

# Dados adicionais posicionados abaixo das faturas
y_dados_adicionais_p2 = y_faturas_p2 - 15*mm - 10*mm
desenhar_dados_adicionais(c, x0=8*mm, y0=y_dados_adicionais_p2, largura_total=194*mm)

# QR Code posicionado abaixo dos dados adicionais
y_qr_code_p2 = y_dados_adicionais_p2 - 25*mm - 10*mm
desenhar_qr_code(c, x0=8*mm, y0=y_qr_code_p2, largura_total=194*mm)

# Salvar
c.save()
print(f"✅ PDF gerado com 2 páginas: {PDF_PATH}")
