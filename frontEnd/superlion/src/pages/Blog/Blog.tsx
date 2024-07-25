import { Box } from "@mui/material";
import BlogCard from "./BlogCard";
import { BlogProps } from "src/models/blog";
import { enqueueSnackbar } from "notistack";
import { getPublicBlogList } from "src/services/blog/blog.service.ts"
import { useEffect, useCallback } from "react";

const Blog = () => {
  const listData: BlogProps[] = [
    
  ];

  const handleGetList = useCallback(async()=>{
    const res = await getPublicBlogList();
    console.log("res",res);
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
        {listData.map((item) => {
          return <BlogCard key={item.id} blog={item} />;
        })}
      </Box>
    </Box>
  );
};

export default Blog;
