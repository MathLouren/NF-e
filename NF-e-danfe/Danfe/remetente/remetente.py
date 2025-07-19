# destinatario.py
from reportlab.lib.units import mm

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_destinatario(c, x0, y0, largura_total):
    altura_total = 28 * mm
    linha_altura = 7 * mm
    col_nome = largura_total * 0.6
    col_cnpj = largura_total * 0.2
    col_data = largura_total * 0.2

    y0 = y0 - 68 * mm  # abaixo do emitente

    # Texto fora da caixa
    c.setFont(FONTE_BOLD, 7)
    c.drawString(x0 + 0*mm, y0 + 1*mm, "DESTINATÁRIO / REMETENTE")

    # Linha 1: NOME/RAZÃO SOCIAL | CNPJ/CPF | DATA
    linha1_y = y0
    linha1_altura = 7 * mm
    c.rect(x0, linha1_y - linha1_altura, largura_total, linha1_altura)
    c.line(x0 + col_nome, linha1_y, x0 + col_nome, linha1_y - linha1_altura)
    c.line(x0 + col_nome + col_cnpj, linha1_y, x0 + col_nome + col_cnpj, linha1_y - linha1_altura)

    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha1_y - 2.5*mm, "NOME/RAZÃO SOCIAL")
    c.drawString(x0 + col_nome + 1.5*mm, linha1_y - 2.5*mm, "C.N.P.J. / C.P.F.")
    c.drawString(x0 + col_nome + col_cnpj + 1.5*mm, linha1_y - 2.5*mm, "DATA DA EMISSÃO")

    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x0 + 1.5*mm, linha1_y - 6*mm, "Matheus Lourenco")
    c.drawString(x0 + col_nome + 1.5*mm, linha1_y - 6*mm, "178.419.912-00")
    c.drawString(x0 + col_nome + col_cnpj + 1.5*mm, linha1_y - 6*mm, "09/02/2025")

    # Linha 2: ENDEREÇO | BAIRRO/DISTRITO | CEP | DATA DE SAÍDA/ENTRADA
    linha2_y = linha1_y - linha1_altura
    col_end = largura_total * 0.5
    col_bairro = largura_total * 0.2
    col_cep = largura_total * 0.10
    col_saida = largura_total * 0.20

    x_end = x0
    x_bairro = x_end + col_end
    x_cep = x_bairro + col_bairro
    x_saida = x_cep + col_cep

    c.rect(x0, linha2_y - linha_altura, largura_total, linha_altura)
    c.line(x_bairro, linha2_y, x_bairro, linha2_y - linha_altura)
    c.line(x_cep, linha2_y, x_cep, linha2_y - linha_altura)
    c.line(x_saida, linha2_y, x_saida, linha2_y - linha_altura)

    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x_end + 1.5*mm, linha2_y - 2.5*mm, "ENDEREÇO")
    c.drawString(x_bairro + 1.5*mm, linha2_y - 2.5*mm, "BAIRRO/DISTRITO")
    c.drawString(x_cep + 1.5*mm, linha2_y - 2.5*mm, "CEP")
    c.drawString(x_saida + 1.5*mm, linha2_y - 2.5*mm, "DATA DE SAÍDA/ENTRADA")

    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x_end + 1.5*mm, linha2_y - 5.5*mm, "Rua das Flores, 832")
    c.drawString(x_bairro + 1.5*mm, linha2_y - 5.5*mm, "Neves")
    c.drawString(x_cep + 1.5*mm, linha2_y - 5.5*mm, "24425-000")
    c.drawString(x_saida + 1.5*mm, linha2_y - 5.5*mm, "09/02/2025")

    # Linha 3: MUNICÍPIO | FONE/FAX | UF | INSCRIÇÃO ESTADUAL | HORA DE SAÍDA
    linha3_y = linha2_y - linha_altura
    col_municipio = largura_total * 0.3
    col_fone = largura_total * 0.2
    col_uf = largura_total * 0.05
    col_insc = largura_total * 0.20
    col_hora = largura_total * 0.25

    x_municipio = x0
    x_fone = x_municipio + col_municipio
    x_uf = x_fone + col_fone
    x_insc = x_uf + col_uf
    x_hora = x_insc + col_insc

    c.rect(x0, linha3_y - linha_altura, largura_total, linha_altura)
    c.line(x_fone, linha3_y, x_fone, linha3_y - linha_altura)
    c.line(x_uf, linha3_y, x_uf, linha3_y - linha_altura)
    c.line(x_insc, linha3_y, x_insc, linha3_y - linha_altura)
    c.line(x_hora, linha3_y, x_hora, linha3_y - linha_altura)

    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x_municipio + 1.5*mm, linha3_y - 2.5*mm, "MUNICÍPIO")
    c.drawString(x_fone + 1.5*mm, linha3_y - 2.5*mm, "FONE/FAX")
    c.drawString(x_uf + 1.5*mm, linha3_y - 2.5*mm, "UF")
    c.drawString(x_insc + 1.5*mm, linha3_y - 2.5*mm, "INSCRIÇÃO ESTADUAL")
    c.drawString(x_hora + 1.5*mm, linha3_y - 2.5*mm, "HORA DE SAÍDA")

    c.setFont(FONTE_TXT, 6.3)
    c.drawString(x_municipio + 1.5*mm, linha3_y - 5.5*mm, "Sao Goncalo")
    c.drawString(x_fone + 1.5*mm, linha3_y - 5.5*mm, "21 98765-4321")
    c.drawString(x_uf + 1.5*mm, linha3_y - 5.5*mm, "RJ")
    c.drawString(x_insc + 1.5*mm, linha3_y - 5.5*mm, "123456789")
    c.drawString(x_hora + 1.5*mm, linha3_y - 5.5*mm, "10:09:37")
    