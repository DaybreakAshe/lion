
import LinearProgress from '@mui/material/LinearProgress';

interface Props {
    loading: boolean;
}

const ProgressBar = (props: Props) => {
    const { loading } = props;
    return (
        <div style={{
            width: "100%",
            height: "4px",
            position: "fixed",
            top: "72px",
            left: 0,
            zIndex: 1,
        }}>
            {loading && <LinearProgress color='success'/>}
        </div>
    )
}
export default ProgressBar;