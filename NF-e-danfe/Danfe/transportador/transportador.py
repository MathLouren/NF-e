# transportador.py
from reportlab.lib.units import mm

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_transportador(c, x0, y0, largura_total, dados_transportador=None):
    """
    Desenha a tabela de TRANSPORTADOR / VOLUMES TRANSPORTADOS no DANFE.
    dados_transportador: dicionário com os campos do transportador (pode ser None para exemplo).
    """
    altura_total = 21 * mm
    linha_altura = 7 * mm
    col_nome = largura_total * 0.30
    col_frete = largura_total * 0.15
    col_antt = largura_total * 0.15
    col_placa = largura_total * 0.15
    col_uf_placa = largura_total * 0.10
    col_cnpj = largura_total - (col_nome + col_frete + col_antt + col_placa + col_uf_placa)

    # Título
    c.setFont(FONTE_BOLD, 7)
    c.drawString(x0 + 0*mm, y0 + 1*mm, "TRANSPORTADOR / VOLUMES TRANSPORTADOS")

    # Linha 1: Nome/Razão Social | Frete por Conta | Código ANTT | Placa do Veículo | UF | CNPJ/CPF
    linha1_y = y0
    c.rect(x0, linha1_y - linha_altura, largura_total, linha_altura)
    x_nome = x0
    x_frete = x_nome + col_nome
    x_antt = x_frete + col_frete
    x_placa = x_antt + col_antt
    x_uf_placa = x_placa + col_placa
    x_cnpj = x_uf_placa + col_uf_placa

    c.line(x_frete, linha1_y, x_frete, linha1_y - linha_altura)
    c.line(x_antt, linha1_y, x_antt, linha1_y - linha_altura)
    c.line(x_placa, linha1_y, x_placa, linha1_y - linha_altura)
    c.line(x_uf_placa, linha1_y, x_uf_placa, linha1_y - linha_altura)
    c.line(x_cnpj, linha1_y, x_cnpj, linha1_y - linha_altura)

    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x_nome + 1.5*mm, linha1_y - 2.5*mm, "NOME/RAZÃO SOCIAL")
    c.drawString(x_frete + 1.5*mm, linha1_y - 2.5*mm, "FRETE POR CONTA")
    c.drawString(x_antt + 1.5*mm, linha1_y - 2.5*mm, "CÓDIGO ANTT")
    c.drawString(x_placa + 1.5*mm, linha1_y - 2.5*mm, "PLACA DO VEÍCULO")
    c.drawString(x_uf_placa + 1.5*mm, linha1_y - 2.5*mm, "UF")
    c.drawString(x_cnpj + 1.5*mm, linha1_y - 2.5*mm, "CNPJ/CPF")

    c.setFont(FONTE_TXT, 6.3)
    if dados_transportador:
        c.drawString(x_nome + 1.5*mm, linha1_y - 6*mm, dados_transportador.get("nome", ""))
        c.drawString(x_frete + 1.5*mm, linha1_y - 6*mm, dados_transportador.get("frete_por_conta", ""))
        c.drawString(x_antt + 1.5*mm, linha1_y - 6*mm, dados_transportador.get("codigo_antt", ""))
        c.drawString(x_placa + 1.5*mm, linha1_y - 6*mm, dados_transportador.get("placa_veiculo", ""))
        c.drawString(x_uf_placa + 1.5*mm, linha1_y - 6*mm, dados_transportador.get("uf_veiculo", ""))
        c.drawString(x_cnpj + 1.5*mm, linha1_y - 6*mm, dados_transportador.get("cnpj", ""))
    else:
        c.drawString(x_nome + 1.5*mm, linha1_y - 6*mm, "EMPRESA EXEMPLO DE TRANSPORTADORAS")
        c.drawString(x_frete + 1.5*mm, linha1_y - 6*mm, "0 - REMETENTE")
        c.drawString(x_antt + 1.5*mm, linha1_y - 6*mm, "")
        c.drawString(x_placa + 1.5*mm, linha1_y - 6*mm, "")
        c.drawString(x_uf_placa + 1.5*mm, linha1_y - 6*mm, "")
        c.drawString(x_cnpj + 1.5*mm, linha1_y - 6*mm, "")

    # Linha 2: Endereço | Município | UF | Inscrição Estadual
    linha2_y = linha1_y - linha_altura
    col_end = largura_total * 0.30
    col_mun = largura_total * 0.20
    col_uf = largura_total * 0.10
    col_ie = largura_total * 0.20
    # espaço restante para a terceira linha
    col_rest = largura_total - (col_end + col_mun + col_uf + col_ie)

    x_end = x0
    x_mun = x_end + col_end
    x_uf = x_mun + col_mun
    x_ie = x_uf + col_uf

    c.rect(x0, linha2_y - linha_altura, largura_total, linha_altura)
    c.line(x_mun, linha2_y, x_mun, linha2_y - linha_altura)
    c.line(x_uf, linha2_y, x_uf, linha2_y - linha_altura)
    c.line(x_ie, linha2_y, x_ie, linha2_y - linha_altura)

    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x_end + 1.5*mm, linha2_y - 2.5*mm, "ENDEREÇO")
    c.drawString(x_mun + 1.5*mm, linha2_y - 2.5*mm, "MUNICÍPIO")
    c.drawString(x_uf + 1.5*mm, linha2_y - 2.5*mm, "UF")
    c.drawString(x_ie + 1.5*mm, linha2_y - 2.5*mm, "INSCRIÇÃO ESTADUAL")

    c.setFont(FONTE_TXT, 6.3)
    if dados_transportador:
        c.drawString(x_end + 1.5*mm, linha2_y - 6*mm, dados_transportador.get("endereco", ""))
        c.drawString(x_mun + 1.5*mm, linha2_y - 6*mm, dados_transportador.get("municipio", ""))
        c.drawString(x_uf + 1.5*mm, linha2_y - 6*mm, dados_transportador.get("uf", ""))
        c.drawString(x_ie + 1.5*mm, linha2_y - 6*mm, dados_transportador.get("inscricao_estadual", ""))
    else:
        c.drawString(x_end + 1.5*mm, linha2_y - 6*mm, "ENDEREÇO EXEMPLO")
        c.drawString(x_mun + 1.5*mm, linha2_y - 6*mm, "IMBITUVA")
        c.drawString(x_uf + 1.5*mm, linha2_y - 6*mm, "PR")
        c.drawString(x_ie + 1.5*mm, linha2_y - 6*mm, "123456789")

    # Linha 3: Quantidade | Espécie | Marca | Numeração | Peso Bruto | Peso Líquido
    linha3_y = linha2_y - linha_altura
    col_qtd = largura_total * 0.15
    col_esp = largura_total * 0.15
    col_marca = largura_total * 0.15
    col_num = largura_total * 0.15
    col_peso_b = largura_total * 0.20
    col_peso_l = largura_total * 0.20

    x_qtd = x0
    x_esp = x_qtd + col_qtd
    x_marca = x_esp + col_esp
    x_num = x_marca + col_marca
    x_peso_b = x_num + col_num
    x_peso_l = x_peso_b + col_peso_b

    c.rect(x0, linha3_y - linha_altura, largura_total, linha_altura)
    c.line(x_esp, linha3_y, x_esp, linha3_y - linha_altura)
    c.line(x_marca, linha3_y, x_marca, linha3_y - linha_altura)
    c.line(x_num, linha3_y, x_num, linha3_y - linha_altura)
    c.line(x_peso_b, linha3_y, x_peso_b, linha3_y - linha_altura)
    c.line(x_peso_l, linha3_y, x_peso_l, linha3_y - linha_altura)

    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x_qtd + 1.5*mm, linha3_y - 2.5*mm, "QUANTIDADE")
    c.drawString(x_esp + 1.5*mm, linha3_y - 2.5*mm, "ESPÉCIE")
    c.drawString(x_marca + 1.5*mm, linha3_y - 2.5*mm, "MARCA")
    c.drawString(x_num + 1.5*mm, linha3_y - 2.5*mm, "NUMERAÇÃO")
    c.drawString(x_peso_b + 1.5*mm, linha3_y - 2.5*mm, "PESO BRUTO")
    c.drawString(x_peso_l + 1.5*mm, linha3_y - 2.5*mm, "PESO LÍQUIDO")

    c.setFont(FONTE_TXT, 6.3)
    if dados_transportador:
        c.drawString(x_qtd + 1.5*mm, linha3_y - 6*mm, dados_transportador.get("quantidade", ""))
        c.drawString(x_esp + 1.5*mm, linha3_y - 6*mm, dados_transportador.get("especie", ""))
        c.drawString(x_marca + 1.5*mm, linha3_y - 6*mm, dados_transportador.get("marca", ""))
        c.drawString(x_num + 1.5*mm, linha3_y - 6*mm, dados_transportador.get("numeracao", ""))
        c.drawString(x_peso_b + 1.5*mm, linha3_y - 6*mm, dados_transportador.get("peso_bruto", ""))
        c.drawString(x_peso_l + 1.5*mm, linha3_y - 6*mm, dados_transportador.get("peso_liquido", ""))
    else:
        c.drawString(x_qtd + 1.5*mm, linha3_y - 6*mm, "1")
        c.drawString(x_esp + 1.5*mm, linha3_y - 6*mm, "CX")
        c.drawString(x_marca + 1.5*mm, linha3_y - 6*mm, "MARCA")
        c.drawString(x_num + 1.5*mm, linha3_y - 6*mm, "123")
        c.drawString(x_peso_b + 1.5*mm, linha3_y - 6*mm, "10,00")
        c.drawString(x_peso_l + 1.5*mm, linha3_y - 6*mm, "9,50")