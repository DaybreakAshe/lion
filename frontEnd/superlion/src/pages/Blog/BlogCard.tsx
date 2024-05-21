import { Box,Typography } from "@mui/material";
import { BlogProps } from "src/models/blog"
import { FC } from "react";

interface Props {
    blog: BlogProps
}


const BlogCard: FC<Props> = ({ blog })=>{
    return (
        <Box sx={{
            width:"100%",
            height:"300px",
            border:"1px solid green"
        }}>
            <Typography>{blog.title}</Typography>
            <Typography>{blog.content}</Typography>
            <Typography>{blog.time}</Typography>
        </Box>
    )
}

export default BlogCard;