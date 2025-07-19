from reportlab.lib.units import mm

FONTE_TXT = "Helvetica"
FONTE_BOLD = "Helvetica-Bold"

def desenhar_dados_adicionais(c, x0, y0, largura_total, dados_adicionais=None):
    """
    Desenha a seção de DADOS ADICIONAIS no DANFE
    dados_adicionais: dicionário com observações do contribuinte e fisco
    """
    altura_total = 25 * mm
    linha_altura = 8 * mm
    
    # Título da seção
    c.setFont(FONTE_BOLD, 7)
    c.drawString(x0 + 0*mm, y0 + 1*mm, "DADOS ADICIONAIS")
    
    # --- Observações do Contribuinte ---
    linha_obs_contribuinte_y = y0
    c.rect(x0, linha_obs_contribuinte_y - linha_altura, largura_total, linha_altura)
    
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha_obs_contribuinte_y - 2.5*mm, "OBSERVAÇÕES DO CONTRIBUINTE")
    
    c.setFont(FONTE_TXT, 6.3)
    if dados_adicionais and dados_adicionais.get("obs_contribuinte"):
        c.drawString(x0 + 1.5*mm, linha_obs_contribuinte_y - 6*mm, dados_adicionais["obs_contribuinte"])
    else:
        c.drawString(x0 + 1.5*mm, linha_obs_contribuinte_y - 6*mm, "Nota fiscal emitida em conformidade com a legislação vigente.")
    
    # --- Observações do Fisco ---
    linha_obs_fisco_y = linha_obs_contribuinte_y - linha_altura
    c.rect(x0, linha_obs_fisco_y - linha_altura, largura_total, linha_altura)
    
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha_obs_fisco_y - 2.5*mm, "OBSERVAÇÕES DO FISCO")
    
    c.setFont(FONTE_TXT, 6.3)
    if dados_adicionais and dados_adicionais.get("obs_fisco"):
        c.drawString(x0 + 1.5*mm, linha_obs_fisco_y - 6*mm, dados_adicionais["obs_fisco"])
    else:
        c.drawString(x0 + 1.5*mm, linha_obs_fisco_y - 6*mm, "Documento autorizado pela SEFAZ.")
    
    # --- Protocolo de Autorização ---
    linha_protocolo_y = linha_obs_fisco_y - linha_altura
    c.rect(x0, linha_protocolo_y - linha_altura, largura_total, linha_altura)
    
    c.setFont(FONTE_BOLD, 6.5)
    c.drawString(x0 + 1.5*mm, linha_protocolo_y - 2.5*mm, "PROTOCOLO DE AUTORIZAÇÃO")
    
    c.setFont(FONTE_TXT, 6.3)
    if dados_adicionais and dados_adicionais.get("protocolo"):
        c.drawString(x0 + 1.5*mm, linha_protocolo_y - 6*mm, dados_adicionais["protocolo"])
    else:
        c.drawString(x0 + 1.5*mm, linha_protocolo_y - 6*mm, "233250053505317 - 09/02/2025 10:09:39")
    
    return y0 - altura_total 