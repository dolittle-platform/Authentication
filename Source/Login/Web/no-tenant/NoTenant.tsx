// Copyright (c) Aigonix. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { Box, Button, Link, Stack, Typography } from '@mui/material';
import { ArrowBack } from '@mui/icons-material';

import AigonixIcon from '../styles/images/aigonix_light_cube.svg';
import DiscordIcon from '../styles/images/discord.svg';
import YoutubeIcon from '../styles/images/youtube.svg';

import { configuration } from '../Configuration';
import { WelcomeHeader } from '../components/WelcomeHeader';

export const NoTenant = () =>
    <Stack>
        <WelcomeHeader />

        <Typography variant='subtitle2' sx={{ mt: 8 }}>
            Explore how to develop with Aigonix across our Developer Channels.
        </Typography>

        <Stack sx={{ mt: 2, gap: 1, justifyItems: 'center', alignSelf: 'center', alignItems: 'flex-start' }}>
            <Button href='https://dolittle.io' target='_blank' startIcon={<AigonixIcon />}>Dolittle SDK Documentation</Button>
            <Button href='https://youtube.com/@Aigonix' target='_blank' startIcon={<YoutubeIcon />}>Aigonix YouTube Channel</Button>
            <Button href='https://discord.gg/tqJQvtw6bY' target='_blank' startIcon={<DiscordIcon />}>Join our Discord</Button>
        </Stack>

        <Box sx={{ mt: 8 }}>
            {configuration.supportEmail &&
                <Typography variant='subtitle2'>
                    An existing organization? <Link href={'mailto:' + configuration.supportEmail}>Contact us</Link> to get started.
                </Typography>
            }
        </Box>

        <Box sx={{ mt: 4 }}>
            <Button color='inherit' startIcon={<ArrowBack />} href={configuration.startPath}>
                Return to login page
            </Button>
        </Box>
    </Stack>;
