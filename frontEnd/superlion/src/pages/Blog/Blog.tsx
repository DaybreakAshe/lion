import { Box } from "@mui/material";
import BlogCard from "./BlogCard";
import { BlogProps } from "src/models/blog";
import { enqueueSnackbar } from "notistack";
import { getPublicBlogList } from "src/services/blog/blog.service.ts"
import { useEffect, useCallback, useState } from "react";

const Blog = () => {
  const [dataList, setDataList] = useState<BlogProps[]>([]);
  const handleGetList = useCallback(async()=>{
    const res = await getPublicBlogList();
    if(res?.code === 200){
      const list = res?.data?.data;
      setDataList(list);
    }else{
      setDataList([]);
    }
  },[]);

  useEffect(() => {
    handleGetList();
  }, [handleGetList]);


  return (
    <Box
      sx={{
        width: "100%",
        display: "flex",
        justifyContent: "center",
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          gap: "15px",
          width: "100%",
        }}
      >
        {dataList?.map((item) => {
          return <BlogCard key={item.id} blog={item} />;
        })}
      </Box>
    </Box>
  );
};

export default Blog;
