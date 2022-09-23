const hexToRgba = (hex: string, opacity: number): string | null => {
  const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex)

  return (result != null) ? `rgba(${parseInt(result[1], 16)}, ${parseInt(result[2], 16)}, ${parseInt(result[3], 16)}, ${opacity})` : null
}

export const getColor = (color: string, opacity: number = 1): string => {
  const c = getComputedStyle(document.body).getPropertyValue(`--tblr-${color}`).trim()
  if (opacity !== 1) {
    return hexToRgba(c, opacity) ?? `rgba(32, 107, 196, ${opacity})`
  }

  return c
}
