import { Box } from "@mui/material";
import BlogCard from "./BlogCard";
import { BlogProps } from "src/models/blog";
import { enqueueSnackbar } from "notistack";

const Blog = () => {
  const listData: BlogProps[] = [
    {
      id: "111",
      title: "是楷体吗",
      content: "content1",
      time: 111,
    },
    {
      id: "222",
      title: "title2",
      content: "content2",
      time: 222,
    },
    {
      id: "333",
      title: "title3",
      content: "content3",
      time: 333,
    },
    {
      id: "444",
      title: "title4",
      content: "content4",
      time: 444,
    },
    {
      id: "555",
      title: "title5",
      content: "content5",
      time: 555,
    },
    {
      id: "666",
      title: "title6",
      content: "content6",
      time: 666,
    },
    {
      id: "777",
      title: "title7",
      content: "content7",
      time: 777,
    },
    {
      id: "888",
      title: "title8",
      content: "content8",
      time: 888,
    },
    {
      id: "999",
      title: "title9",
      content: "content9",
      time: 999,
    },
  ];

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
