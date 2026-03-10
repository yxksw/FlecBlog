// 代理后端的 Atom Feed
export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig()
  // 从 apiUrl 提取后端基础地址（移除 /api/vX 部分）
  const backendUrl = config.public.apiUrl.replace(/\/api\/v\d+$/, '')
  
  try {
    // 直接使用原生 fetch 获取完整响应
    const response = await fetch(`${backendUrl}/atom.xml`)
    
    if (!response.ok) {
      throw new Error(`Backend returned ${response.status}`)
    }
    
    // 获取原始文本内容
    const atomFeed = await response.text()
    
    // 复制后端的响应头
    const contentType = response.headers.get('content-type') || 'application/atom+xml; charset=utf-8'
    setResponseHeader(event, 'Content-Type', contentType)
    
    // 返回原始 XML 内容（包含 XML 声明）
    return atomFeed
  } catch (error) {
    throw createError({
      statusCode: 500,
      statusMessage: 'Failed to fetch Atom feed'
    })
  }
})
