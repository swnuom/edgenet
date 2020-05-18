import React from "react";
import { Box } from "grommet";

const GrammaireList = ({item, onClick}) =>
    <Box pad="small" onClick={() => onClick(item.id)}>
        {item.title}
    </Box>;

export default GrammaireList;
