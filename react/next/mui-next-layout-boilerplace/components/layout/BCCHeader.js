import {
    AppBar,
    Container,
    Grid,
    Toolbar,
} from '@mui/material'
import Image from "next/image";
import Link from 'next/link'

const BCCHeader = () => {
    return (
        <AppBar variant="dense" color="transparent">
            <Container sx={{mt: 2}}>
                <Toolbar variant="dense">
                    <Grid
                        container
                        direction="row"
                        justifyContent={{ xs: 'center', sm: 'center', md: 'space-between' }}
                        alignItems="center"
                    >
                            <Link href='/'>
                                <Image src='' width='75px' height='50px' />
                            </Link>
                    </Grid>
                </Toolbar>
            </Container>
        </AppBar>
    )
}

export default BCCHeader
