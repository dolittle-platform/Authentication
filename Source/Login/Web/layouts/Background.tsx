// Copyright (c) Dolittle. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

import React from 'react';

import { configuration } from '../Configuration';

const Hexagon = (props: React.SVGProps<SVGPathElement>): JSX.Element =>
    <g transform={props.transform}>
        <path 
            fill={props.fill}
            filter={props.filter}
            transform="scale(0.0425,0.0425) translate(-39.2,-40.2)"
            d="M58.14,40.22H46.9a2,2,0,0,0-1.69,1l-5.62,9.74a1.94,1.94,0,0,0,0,2l5.62,9.74a2,2,0,0,0,1.69,1H58.14a2,2,0,0,0,1.69-1l5.62-9.74a1.94,1.94,0,0,0,0-2l-5.62-9.74a2,2,0,0,0-1.69-1Z"
            />
    </g>

const Logo = (props: React.SVGProps<SVGGElement>): JSX.Element =>
    <g transform={props.transform}>
        <path 
            fill={props.fill}
            filter={props.filter}
            transform="scale(0.013,0.013) translate(-31,-14)"
            d="M58.14,40.22H46.9a2,2,0,0,0-1.69,1l-5.62,9.74a1.94,1.94,0,0,0,0,2l5.62,9.74a2,2,0,0,0,1.69,1H58.14a2,2,0,0,0,1.69-1l5.62-9.74a1.94,1.94,0,0,0,0-2l-5.62-9.74a2,2,0,0,0-1.69-1Z"
            />
        <path
            fill={props.fill}
            filter={props.filter}
            transform="scale(0.013,0.013) translate(-31,-14)"
            d="M75.84,14.39a1,1,0,0,0-1.66,0L65.09,30.14a1.6,1.6,0,0,0,0,1.61L76.37,51.3a1.19,1.19,0,0,1,0,1.21L65,72.26a1.19,1.19,0,0,1-1,.6H41.35a1.63,1.63,0,0,0-1.4.81L30.86,89.42a1,1,0,0,0,.83,1.43H72.51a4.31,4.31,0,0,0,3.74-2.15l20-34.63a4.33,4.33,0,0,0,0-4.32Z"
        />
    </g>;

export const Background = (): JSX.Element => {
    return (
        <svg width='100%' height='100%' viewBox='0 0 2880 1800' preserveAspectRatio='xMinYMax slice' css={{ position: 'absolute' }}>
            <defs>
                <filter id='blur-8' x='-100%' y='-100%' width='300%' height='300%'>
                    <feGaussianBlur stdDeviation='8' />
                </filter>
                <filter id='blur-12' x='-150%' y='-150%' width='400%' height='400%'>
                    <feGaussianBlur stdDeviation='12' />
                </filter>
                <filter id='blur-14' x='-150%' y='-150%' width='400%' height='400%'>
                    <feGaussianBlur stdDeviation='14' />
                </filter>
                <filter id='blur-18' x='-150%' y='-150%' width='400%' height='400%'>
                    <feGaussianBlur stdDeviation='18' />
                </filter>
            </defs>
            <g>
                <Logo filter='url(#blur-8)' transform='translate(540, 992) scale(-731.4,-731.4)' fill='#76E8DB4D'/>
                <Hexagon filter='url(#blur-12)' transform='translate(413,228) scale(330,330)' fill='#FF936666' />
                { configuration.animateBackground && <>
                    <animateTransform
                        attributeName='transform' attributeType='XML' type='translate'
                        values='0 0; -800 -800'
                        keyTimes='0; 1'
                        keySplines='0.5 0 1 0.5'
                        dur='60s'
                        calcMode='spline' repeatCount='1' fill='freeze' />
                    <animateTransform
                        attributeName='transform' attributeType='XML' type='scale'
                        values='1 1; 2 2'
                        keyTimes='0; 1'
                        keySplines='0.5 0 1 0.5'
                        dur='60s'
                        calcMode='spline' repeatCount='1' fill='freeze' additive='sum' />
                </>}
            </g>
            <g transform='translate(-85,1242.43)'>
                <Logo filter='url(#blur-8)' transform='scale(731.4,731.4)' fill='#FFB79966'/>
                <Hexagon filter='url(#blur-8)' transform='translate(60,230) scale(254,254)' fill='#FFB799'/>
            </g>
            <Hexagon filter='url(#blur-18)' transform='translate(153,530) scale(330,330)' fill='#48E0CF' />
            <g>
                <Hexagon filter='url(#blur-12)' transform='translate(718,1084) scale(228,228)' fill='#6678F6' />
                <Hexagon filter='url(#blur-12)' transform='translate(720,1452) scale(228,228)' fill='#6678F680' />
                <Hexagon filter='url(#blur-14)' transform='translate(897,1231) scale(329,329)' fill='#76E8DB' />
                { configuration.animateBackground && <>
                    <animateTransform
                        attributeName='transform' attributeType='XML' type='translate'
                        values='0 0; 3000 0'
                        keyTimes='0; 1'
                        keySplines='0.5 0 1 0.5'
                        dur='300s'
                        calcMode='spline' repeatCount='1' fill='freeze' />
                    <animateTransform
                        attributeName='transform' attributeType='XML' type='scale'
                        values='1 1; 2 2'
                        keyTimes='0; 1'
                        keySplines='0.5 0 1 0.5'
                        dur='300s'
                        calcMode='spline' repeatCount='1' fill='freeze' additive='sum' />
                </>}
            </g>
            <g>
                <Hexagon filter='url(#blur-8)' transform='translate(1418,-600) scale(330,330)' fill='#6678F6' />
                <Hexagon filter='url(#blur-12)' transform='translate(1518,-1000) scale(570,570)' fill='#76E8DB' />
                { configuration.animateBackground && <>
                    <animateTransform
                        attributeName='transform' attributeType='XML' type='translate'
                        values='0 0; 0 0; -800 4000'
                        keyTimes='0; 0.5; 1'
                        dur='300s'
                        repeatCount='indefinite'/>
                    <animateTransform
                        attributeName='transform' attributeType='XML' type='scale'
                        values='1 1; 0.6 0.6'
                        keyTimes='0; 1'
                        dur='300s'
                        repeatCount='indefinite' additive='sum' />
                </>}
            </g>
        </svg>
    );
};