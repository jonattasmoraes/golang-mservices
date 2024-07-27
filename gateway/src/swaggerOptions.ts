import swaggerJSDoc, { Options } from 'swagger-jsdoc'

const swaggerOptions: Options = {
    definition: {
        openapi: '3.0.0',
        info: {
            title: 'Microservice Gateway API',
            version: '1.0.0',
            description: 'Sale microservice gateway API Documentation',
        },
        components: {
            schemas: {
                ErrorResponse: {
                    type: 'object',
                    properties: {
                        data: {
                            type: 'object',
                            properties: {
                                error: {
                                    type: 'string',
                                    example:
                                        'An error occurred while processing the request. Please try again later.',
                                },
                            },
                        },
                    },
                },
                SaleResponse: {
                    type: 'object',
                    properties: {
                        data: {
                            type: 'object',
                            properties: {
                                UserID: {
                                    type: 'string',
                                    example: '01J1P83E4KJ3EM3H73139AC5PQ',
                                },
                                SaleID: {
                                    type: 'string',
                                    example: '01J1P83E4KJWRSCDHHJSKKDJDG',
                                },
                                PaymentType: {
                                    type: 'string',
                                    example: 'credit',
                                },
                                Date: {
                                    type: 'string',
                                    example: '2022-01-01T00:00:00.000Z',
                                },
                                Products: {
                                    type: 'array',
                                    items: {
                                        type: 'object',
                                        properties: {
                                            ProductID: {
                                                type: 'string',
                                                example: '01J1PATMVDPDQJKRC037EGDFET',
                                            },
                                            Name: {
                                                type: 'string',
                                                example: 'refrigerante coca-cola 350ml',
                                            },
                                            Unit: {
                                                type: 'string',
                                                example: 'un',
                                            },
                                            Category: {
                                                type: 'string',
                                                example: 'bebidas',
                                            },
                                            Quantity: {
                                                type: 'integer',
                                                example: 2,
                                            },
                                            Price: {
                                                type: 'number',
                                                example: 5,
                                            },
                                        },
                                    },
                                },
                                Total: {
                                    type: 'number',
                                    example: 10,
                                },
                            },
                        },
                        message: {
                            type: 'string',
                            example: 'Sale created successfully',
                        },
                    },
                },
            },
        },
    },
    apis: ['./src/**/*.ts'], // Ajuste para incluir os arquivos corretos
}

const swaggerSpec = swaggerJSDoc(swaggerOptions)

export default swaggerSpec
