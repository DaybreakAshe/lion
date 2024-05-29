import { useEffect, useCallback } from "react";
import { getMyBlogList } from "src/services/blog/blog.service.ts";

const MyBlog = () => {
  const handleList = useCallback(async () => {
    const param = {
      id: "11",
      title: "",
    };
    const res = await getMyBlogList(param);
    console.log("list", res);
  }, []);
  useEffect(() => {
    handleList();
  }, [handleList]);
  return <></>;
};
export default MyBlog;
