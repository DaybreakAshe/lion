import { Box, Theme } from "@mui/material";
import { makeStyles } from '@mui/styles'
import { useState } from "react";
import { Editor } from '@tinymce/tinymce-react';
import ConfirmButton from "../../components/ConfirmButton/ConfirmButton";
const useStyles = makeStyles((_theme: Theme) => ({
    root: {
        width: "100%",
        height: '500px',
    },
    buttonBox:{
        display:"flex",
        justifyContent:"flex-end",
        marginTop:"20px",
    }
}))
const PublicBlog = () => {
    const classes = useStyles()
    const [content, setContent] = useState<any>(null);
    const saveBlog = () => {
        console.log("###",content,typeof content)
        //转json，传给后端
        const contentJson = JSON.stringify(content)
        console.log("###",contentJson,typeof contentJson)
    }
    const cancel = () => {
        console.log("cancel")
    }
    return (
        <>
            <Box className={classes.root}>
                <Editor
                    apiKey='hxv1yh7415bxinnjobb42way775jlu1rlil885zbyk4fw7om'
                    init={{
                        height: 500,
                        menubar: true,
                        plugins: [
                            'advlist', 'autolink', 'lists', 'link', 'image', 'charmap', 'preview',
                            'anchor', 'searchreplace', 'visualblocks', 'code', 'fullscreen',
                            'insertdatetime', 'media', 'table', 'code', 'help', 'wordcount'
                        ],
                        toolbar: 'undo redo | blocks | ' +
                            'bold italic forecolor | alignleft aligncenter ' +
                            'alignright alignjustify | bullist numlist outdent indent | ' +
                            'removeformat | link | help',
                        content_style: 'body { font-family:Helvetica,Arial,sans-serif; font-size:14px }'
                    }}
                    onEditorChange={(newValue) => {
                        setContent(newValue);
                    }}
                    value={content}
                />
                <Box className={classes.buttonBox}>
                    <ConfirmButton
                        loading={false}
                        value={"Cancel"}
                        handleClick={cancel}
                        option={{
                            width: "100px",
                            height: "42px",
                            background: "#fff",
                            marginRight: "20px",
                        }}
                    />
                    <ConfirmButton
                        loading={false}
                        value={"Save"}
                        handleClick={saveBlog}
                        option={{
                            width: "100px",
                            height: "42px",
                            color: "#fff",
                        }}
                    />
                </Box>
                
            </Box>
        </>
    )
}

export default PublicBlog;