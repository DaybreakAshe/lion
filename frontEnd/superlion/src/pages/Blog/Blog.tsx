import { Box } from "@mui/material";
import { makeStyles } from '@mui/styles'
import BlogCard from "./BlogCard"
import { BlogProps } from "src/models/blog"
const useStyles = makeStyles(() => ({
    root: {
        
    },
}))

const Blog = () => {
    const classes = useStyles()
    const listData: BlogProps[] = [
        {
            id:"111",
            title:"title1",
            content:"content1",
            time:111
        },
        {
            id: "222",
            title:"title2",
            content:"content2",
            time:222
        },
        {
            id:"333",
            title:"title3",
            content:"content3",
            time:333
        }
    ]

    return (
        <>
            <Box sx={{
                border: "1px solid red",
                display: "flex",
                flexDirection: "column",
                gap: "15px"
            }}>
                {listData.map((item)=>{
                    return <BlogCard key={item.id} blog={item}/>
                })}
            </Box>
        </>
    )
}

export default Blog;