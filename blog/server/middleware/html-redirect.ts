import { defineEventHandler, sendRedirect } from 'h3'

/**
 * 伪静态路由中间件
 * 将 /posts/xxx.html 重定向到 /posts/xxx
 */
export default defineEventHandler((event) => {
  const url = event.path
  
  // 匹配 /posts/xxx.html 格式的 URL
  const match = url.match(/^\/posts\/(.+)\.html$/)
  
  if (match) {
    const slug = match[1]
    // 301 永久重定向到不带 .html 的 URL
    return sendRedirect(event, `/posts/${slug}`, 301)
  }
})
