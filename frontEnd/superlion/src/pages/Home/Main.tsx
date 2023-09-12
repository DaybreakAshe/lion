import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import homeImage from "../../../src/assets/images/home/index.jpg"
import ArticleCard from "./ArticleCard";
import { uuid } from "../../utils/helpers";
import { ArticleCardProps } from "@/services/home/article.service";
import { useState, useEffect } from "react";
import { getArticleList } from "../../services/home/article.service";
const useStyles = makeStyles((theme: Theme) => ({
    content: {
        width: "100%",
        position: "absolute",
        top: "65px",
        boxSizing: "border-box",
    },
    articleBox: {
        maxWidth: "1200px",
        margin: "0 auto",
        padding: "20px",
        paddingTop: "0",
    }
}))

const articleList = [
    {
        title: "test",
        content: "Death investigated at Burning Man while 70,000 festival attendees remain stuck in Nevada desert after rain",
        image: homeImage,
        date: "2021-10-10",
    },
    {
        title: "test",
        content: "Authorities are investigating a death at the Burning Man festival in the Nevada desert as thousands of people remain trapped on site after heavy rains inundated the area and created thick, ankle-deep mud which sticks to campers’ shoes and vehicle tires.Authorities are investigating a death at the Burning Man festival in the Nevada desert as thousands of people remain trapped on site after heavy rains inundated the area and created thick, ankle-deep mud which sticks to campers’ shoes and vehicle tires.Authorities are investigating a death at the Burning Man festival in the Nevada desert as thousands of people remain trapped on site after heavy rains inundated the area and created thick, ankle-deep mud which sticks to campers’ shoes and vehicle tires.Authorities",
        image: homeImage,
        date: "2021-10-10",
    },
    {
        title: "test",
        content: "test",
        image: homeImage,
        date: "2021-10-10",
    },
    {
        title: "test",
        content: "test",
        image: homeImage,
        date: "2021-10-10",
    }
]

const Main = () => {
    const classes = useStyles()
    const [list, setList] = useState<ArticleCardProps[]>([])
    const getArticle = async () => {
        const res = await getArticleList()
        console.log("res", res)
    }
    useEffect(() => {
        getArticle()
    }, [])
    return (
        <>
            <Box className={classes.content}>
                <Box className={classes.articleBox}>
                    {
                        list.map((item, _index) => (
                            <Box key={uuid()}>
                                <ArticleCard title={item.title} content={item.content} image={item.image} date={item.date} />
                            </Box>
                        ))
                    }
                </Box>
            </Box>
        </>
    )
}

export default Main