import { Box, Theme, TextField } from "@mui/material";
import { makeStyles } from '@mui/styles'
import { useState } from "react";
const useStyles = makeStyles((theme: Theme) => ({
    content: {
        width: "100%",
        padding:'20px',
        margin:"20px 0",
        borderRadius:'8px',
        background:'#fff',
        boxSizing:'border-box',
    },
    item:{
        display:'flex',
        alignItems:'center',
        marginBottom:'20px',
    },
    title:{
        display:'inline-block',
        width:'100px',
        textAlign:'center',
        marginRight:'10px',
        fontSize:'18px',
        fontWeight:'bold',
        fontFamily:'PingFangSC-Regular',
        [`${theme.breakpoints.down('sm')}`]: {
            fontSize:'14px',
            width:'80px',
        }
    },
    inputStyle:{
        width:'300px',
        "& .MuiInputBase-root": {
            borderRadius:"8px",
        },
    }
}))

const SetBlogInfo = () => {
    const classes = useStyles()
    const [TitleValue, setCampaignName] = useState<string>(''); //标题
    return (
        <>
            <Box className={classes.content}>
                <Box className={classes.item}>
                    <span className={classes.title}>Title</span>
                    <TextField
                        variant="outlined"
                        placeholder="Enter Title"
                        type='text'
                        value={TitleValue}
                        onChange={(e) => { setCampaignName(e.target.value) }}
                        className={classes.inputStyle}
                    />
               </Box>
            </Box>
        </>
    )
}

export default SetBlogInfo;

