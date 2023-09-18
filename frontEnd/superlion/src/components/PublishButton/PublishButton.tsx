import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';
const useStyles = makeStyles((theme: Theme) => ({
    root: {

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
    }
}))

const PublishButton = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
                <Box className={classes.publishButton}>
                    <span>Publish</span>
                    <ExpandMoreIcon style={{
                        fontSize: "20px",
                    }}/>
                </Box>
            </Box>
        </>
    )
}

export default PublishButton;