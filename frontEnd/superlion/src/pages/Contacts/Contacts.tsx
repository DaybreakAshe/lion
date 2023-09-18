import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
const useStyles = makeStyles((_theme: Theme) => ({
    root: {
    },
}))

const Contacts = () => {
    const classes = useStyles()
    return (
        <>
            <Box className={classes.root}>
                Contacts
            </Box>
        </>
    )
}

export default Contacts;