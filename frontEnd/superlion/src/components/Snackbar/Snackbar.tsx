import React, { FC, useState } from 'react';
import { Snackbar } from '@mui/material';
import MuiAlert, { AlertProps } from '@mui/material/Alert';
const Alert: FC<AlertProps> = (props) => {
    return <MuiAlert elevation={6} variant="filled" {...props} />;
};
export interface Props {
    isOpen: boolean;
    message: string;
    severity?: 'error' | 'warning' | 'info' | 'success';
    duration?: number;
    vertical?: 'top' | 'bottom';
    horizontal?: 'left' | 'right' | 'center';
    key?: number;
}
const SnackbarMessage: FC<Props> = ({ message, severity = 'info', duration = 15000, isOpen = false, vertical = 'bottom', horizontal = 'left', key }) => {
    const [open, setOpen] = useState<boolean>(false);
    const handleClose = () => {
        setOpen(false);
    };
    React.useEffect(() => {
        if (message) {
            setOpen(true);
        }
    }, [message, isOpen]);
    return (
        <Snackbar
            open={open}
            autoHideDuration={duration}
            onClose={() => handleClose()}
            anchorOrigin={{ vertical, horizontal }}
            key={key}
        >
            <div>
                <Alert onClose={handleClose} severity={severity}>
                    {message}
                </Alert>
            </div>
        </Snackbar>
    );
};

export default SnackbarMessage;
