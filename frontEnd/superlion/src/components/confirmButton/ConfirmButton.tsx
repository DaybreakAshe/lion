import LoadingButton from '@mui/lab/LoadingButton';
interface loadingButton {
    loading: boolean;
    isShow?: boolean,
    value: string;
    handleClick: () => any,
    option?: {
        icon?: string,
        width?: string,
        height?: string,
        background?: string,
        boxShadow?: string,
        borderRadius?: string,
        color?: string,
        marginLeft?: string,
        border?: string,
        textTransform?: string,
        marginTop?: string,
        fontSize?: string,
        backgroundHover?: string,
        borderColorHover?: string,
        marginRight?: string,
        marginBottom?: string,
        isDisabled?: boolean,
        isPointerEvents?: 'none' | 'auto' | 'inherit',
    }
}

const LoadingBTN = (props: loadingButton) => {
    const { loading, value, handleClick, option, isShow = true } = props
    return (
        <>
            {isShow && <LoadingButton
                disabled={loading || option?.isDisabled}
                variant="outlined"
                loading={loading}
                onClick={(e) => {
                    e.stopPropagation();
                    e.nativeEvent.stopImmediatePropagation();
                    handleClick()
                }}
                style={{
                    background: loading ? '#fff' : option?.background || "#fff !important",
                    pointerEvents: loading ? 'none' : (option?.isPointerEvents || 'auto'),
                }}
                sx={{
                    width: option?.width || "109px !important",
                    height: option?.height || "36px !important",
                    boxShadow: option?.boxShadow || "0px 2px 5px rgba(48, 125, 207, 0.1) !important",
                    borderRadius: option?.borderRadius || "8px !important",
                    marginLeft: option?.marginLeft || "0px",
                    marginRight: option?.marginRight || "0px",
                    textTransform: 'none',
                    marginTop: option?.marginTop || "0 !important",
                    marginBottom: option?.marginBottom || "0 !important",
                    fontSize: option?.fontSize || "13px",
                    background: option?.background || "#1a73e8",
                    color: option?.color || "#1a73e8",
                    border: option?.border || `1px solid #1a73e8 !important`,
                    "&:hover": {
                        borderColor: option?.borderColorHover || "#FB6B5A",
                        background: option?.backgroundHover || "#1a73e8"
                    },
                }}>
                {option?.icon && <img src={option.icon} alt="" style={{ margin: "5px", width: "17px", height: "17px" }} />}
                {!loading && value}
            </LoadingButton>}
        </>
    );
}
export default LoadingBTN