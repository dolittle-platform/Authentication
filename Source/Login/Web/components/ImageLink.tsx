// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import { ReactNode } from 'react';

import { Box, Link, LinkProps } from '@mui/material';

export type ImageLinkProps = {
    href: string;
    image?: ReactNode;
    text: string
    target?: string;
    linkProps?: LinkProps;
}

export const ImageLink = ({ href, image, text, target = '_blank', linkProps }: ImageLinkProps) => (
    <Link
        href={href}
        target={target}
        {...linkProps}
        sx={{
            display: 'flex',
            gap: 2,
            justifyContent: 'center',
        }}
    >
        <Box sx={{ width: 34, verticalAlign:'center' }} display='flex' justifyContent='center'>{image}</Box>
        <Box component='span' display='flex' justifyContent='center'>{text}</Box>
    </Link>
);