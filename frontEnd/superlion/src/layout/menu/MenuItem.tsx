import { Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import { Link } from "react-router-dom";
import { useLocation } from "react-router-dom";
const useStyles = makeStyles((_theme: Theme) => ({
    content: {

    },
    item: {
        color: "black",
        fontSize: "18px",
        display: "inline-block",
    }
}))

interface MenuItemProps {
    title: string,
    path: string,
    isMobile?: boolean,
}

const MenuItem = (props: MenuItemProps) => {
    const { title, path, isMobile = false } = props
    const classes = useStyles()
    const location = useLocation();
    const pathSegments = location.pathname.split('/');
    const lastSegment = pathSegments[pathSegments.length - 1];
    return (
        <>
            <Link to={path} style={{
                textDecoration: "none",
                width: isMobile ? "100%" : "auto",
            }}>
                <span
                    className={classes.item}
                    style={{
                        textAlign: isMobile ? "center" : "left",
                        marginRight: isMobile ? "0px" : "30px",
                        height: isMobile ? "40px" : "auto",
                        lineHeight: isMobile ? "40px" : "auto",
                        width: isMobile ? "100%" : "auto",
                        color: lastSegment === path.split("/")[1] ? "#1a73e8" : "black",
                    }}
                >{title}</span>
            </Link>
        </>
    )
}

export default MenuItem;