import React from "react";
import { Box, Text } from "grommet";

const NewsList = ({item, onClick}) =>
    <Box pad="small" onClick={() => onClick(item.id)}>
        <Text size="small">
            {item.date}
        </Text>
        {item.text.substring(0,150)}...
    </Box>;

export default NewsList;
