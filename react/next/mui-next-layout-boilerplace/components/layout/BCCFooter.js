import {
    Box,
    Container,
    Grid,
    IconButton,
    styled, SvgIcon,
    Toolbar,
} from '@mui/material'
import { mdiTwitter } from '@mdi/js'
import React from 'react'

const BCCWhiteSvgIcon = styled(SvgIcon)({
    color: "#ffffff"
})

export function BCCFooter() {
    return (
        <footer>
            <Box mt={2}>
                <Container>
                    <Toolbar>
                        <Grid
                            container
                            fullWidth
                            direction="row"
                            alignItems="center"
                            justifyContent={{ xs: 'center', md: 'flex-end'}}
                        >
                            <IconButton
                                href="/"
                                target="_blank"
                            >
                                <BCCWhiteSvgIcon color='success'>
                                    <path d={mdiTwitter}/>
                                </BCCWhiteSvgIcon>
                            </IconButton>
                        </Grid>
                    </Toolbar>
                </Container>
            </Box>
        </footer>
    )
}
