export interface Tag {
  TagId: string;
  Tag: string;
}

export interface BlogProps {
  id: number;
  title: string;
  headImg: string;
  official: number; // 官方1-是，0-否
  auditState: string;
  views: number;
  approvals: number; // 点赞量
  collection: number; // 收藏量
  sort: number;
  authorId: string;
  preview: string; // 预览内容
  tags: Tag[]; // 根据你的 model.Tag 结构体定义相应的字段
}

export interface GetListResponse {
  data: BlogProps[];
  code: number;
  msg: string;
  total: number;
}
