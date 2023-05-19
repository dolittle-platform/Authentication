// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.


import DiscordLogo from '../styles/images/discord.svg';
import YoutubeLogo from '../styles/images/youtube.svg';
import AigonixCube from '../styles/images/aigonix_cube_dark.svg';


import { Box, Link, Typography } from '@mui/material';

import { configuration } from '../Configuration';
import { WelcomeHeader } from '../components/WelcomeHeader';
import { ImageLink } from '../components/ImageLink';

export const NoTenant = (): JSX.Element => {
    const { applicationName, logoutPath, supportEmail } = configuration;

    return (
        <Box>
            <WelcomeHeader />

            <Typography variant='subtitle2'>
                Thanks for your interest in {applicationName}!
            </Typography>
            <Typography variant='subtitle2' sx={{ mt: 3 }}>
                Our Studio will go live soon.
            </Typography>
            <Typography variant='subtitle2' sx={{ mt: 3 }}>
                In the meantime, please explore how to develop with Aigonix across our Developer Channels.
            </Typography>

            <Box display='flex' justifyContent='center' mt={5}>
                <Box display='flex' flexDirection='column' gap={3} alignItems='flex-start' justifyItems='center'>
                    <ImageLink href='https://dolittle.io' text='Dolittle SDK Documentation' image={<AigonixCube />} />
                    <ImageLink href='https://youtube.com/@Aigonix' image={<YoutubeLogo />} text='Aigonix YouTube Channel' />
                    <ImageLink href='https://discord.gg/tqJQvtw6bY' image={<DiscordLogo />} text='Join our Discord!' />
                </Box>
            </Box>

            <Box sx={{ mt: 12.5, mb: 5 }}>
                {
                    supportEmail &&
                    <Typography variant='subtitle2'>
                        An existing customer? <Link href={'mailto:' + supportEmail}>Contact us</Link> to get started.
                    </Typography>
                }
            </Box>
        </Box>
    );
};