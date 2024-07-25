import { Box, Divider } from "@mui/material";
import { BlogProps } from "src/models/blog";
import { FC } from "react";

interface Props {
  blog: BlogProps;
}

const BlogCard: FC<Props> = ({ blog }) => {
  return (
    <Box
      sx={{
        width: "100%",
        padding: "10px 20px",
        backgroundColor:"#fff",
        boxShadow:"0 0 10px rgba(0,0,0,.1)",
        borderRadius:"5px",
        boxSizing:"border-box"
      }}
    >
      <Box
        sx={{
          height: "60px",
          display: "flex",
          justifyContent: "left",
          alignItems: "center",
          fontSize:"20px",
          fontWeight:600
        }}
      >
        {blog.title}
      </Box>
      <Divider />
      <Box
        sx={{
          height: "200px",
          padding:"15px 0"
        }}
      >
        {blog.content}
      </Box>
      <Divider />
      <Box
        sx={{
          height: "40px",
          display: "flex",
          justifyContent: "right",
          alignItems: "center",
          fontSize:"17px",
          fontWeight:500,
          color:"#666"
        }}
      >
        {blog.time}
      </Box>
    </Box>
  );
};

export default BlogCard;
