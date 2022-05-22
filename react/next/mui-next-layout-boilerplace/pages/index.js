import BCCLayout from "../components/layout/BCCLayout";
import {Stack, Typography} from "@mui/material";
import '@fontsource/permanent-marker'
import '@fontsource/karla'

const Home = () => {
  return (
      <BCCLayout>
        <Stack alignItems="center" justifyContent='center' spacing={2} mt={8} style={{ height: '100%'}}>
            <Typography>Hello, Friend</Typography>
        </Stack>
      </BCCLayout>
  )
}

export default Home
