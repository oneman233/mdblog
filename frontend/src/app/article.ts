// 文章类的通用接口
export interface Article {
  id?: number; // 声明可选字段
  title?: string
  time?: string
  writer?: string
  abstract?: string
  content?: string
}
