import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import { ArticleCardProps } from "@/services/home/article.service";
const useStyles = makeStyles((theme: Theme) => ({
    content: {
        width: "100%",
        height: "300px",
        backgroundColor: "#fff",
        borderRadius: "10px",
        boxShadow: "0 0 10px rgba(0,0,0,.1)",
        padding: "20px",
        boxSizing: "border-box",
        margin: "20px 0",
        display: "flex",
        justifyContent: "space-between",
    },
    title: {
        fontSize: "20px",
        fontWeight: "bold",
        height: "40px",
        
    },
    textBox: {
        width: "70%",
        height: "100%", 
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        boxSizing: "border-box",
        padding: "20px",
    },
    imageBox: {
        width: "30%",
        height: "100%",
    },
    time:{
        fontSize: "12px",
        color: "#999",
        width: "100%",
        textAlign: "right",
    },
    textContent:{
        fontSize: "14px",
        width: "100%",
        height: "100%",
        overflow: "hidden",
        textOverflow: "ellipsis",
        display: "-webkit-box",
        "-webkit-line-clamp": "3",
        "-webkit-box-orient": "vertical",
        lineHeight: "20px",
    }
}))

const ArticleCard = (props: ArticleCardProps) => {
    const { title, content, image, date } = props
    const classes = useStyles()
    return (
        <>
            <Box className={classes.content}>
                <Box className={classes.imageBox} style={{
                    backgroundImage: `url(${image})`,
                    backgroundSize: "cover",
                    backgroundPosition: "center",
                }}></Box>
                <Box className={classes.textBox}>
                    <Box className={classes.title}>{title}</Box>
                    <Box className={classes.textContent}>{content}</Box>
                    <Box className={classes.time}>{date}</Box>
                </Box>
            </Box>
        </>
    )
}

export default ArticleCard