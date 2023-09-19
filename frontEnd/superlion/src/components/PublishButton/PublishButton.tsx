import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
import { useState, useEffect } from "react";
import MenuItem from "../../layout/Menu/MenuItem";
const useStyles = makeStyles((theme: Theme) => ({
    root: {
        position: "relative",
    },
    publishButton: {
        background: "#1a73e8",
        color: "#fff",
        width: "85px",
        height: "40px",
        borderRadius: "8px",
        marginRight: "30px",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        fontSize: "14px",
        cursor: "pointer",
    },
    publishCard: {
        width: "130px",
        position: "absolute",
        right: "20px",
        background: "#fff",
        top: "40px",
        borderRadius: "8px",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        boxShadow: "0px 0px 10px rgba(0,0,0,0.1)",
        transition: "all 0.3s ease",
        overflow: "hidden",
    }
}))

const PublishButton = () => {
    const classes = useStyles()
    const [open, setOpen] = useState(false);
    useEffect(() => {
        document.addEventListener('click', () => {
            setOpen(false)
        })
    }, [])
    return (
        <>
            <Box className={classes.root}>
                <Box className={classes.publishButton} onClick={(e: any) => {
                    e.stopPropagation();
                    setOpen(!open)
                }}>
                    <span>Publish</span>
                    <ExpandMoreIcon style={{
                        fontSize: "20px",
                        transform: open ? "rotate(180deg)" : "rotate(0deg)",
                    }} />
                </Box>
                <Box className={classes.publishCard} style={{
                    height: open ? "120px" : "0px",
                    padding: open ? "10px 0" : "0px",

                }}>
                    <MenuItem title="Publish blog" path="/publishBlog" isMobile />
                    <MenuItem title="My Blog" path="/myBlog" isMobile />
                    <MenuItem title="My Draft" path="/myDraft" isMobile />
                </Box>
            </Box>
        </>
    )
}

export default PublishButton;