import * as Koa from 'koa';
import {ApolloServer, gql} from 'apollo-server-koa';
import * as pubs from './papers.json';
import * as loki from 'lokijs';
import {IPubs} from './ipubs';
import { publicDecrypt } from 'crypto';

var db = new loki('my-app')
let pubsC = db.addCollection('pubs')
pubs.map(p => pubsC.insert(p))

// Construct a schema, using GraphQL schema language
const typeDefs = gql`
  type PubResponse {
    id : Int
    title: String
    cite: String
    link: String
    slides: String
    abstract: String
  }
  type Query {
    pub( pid: Int): PubResponse
    pubs: [PubResponse]
    pquery (qry: String): [PubResponse]
  }
`;

// Provide resolver functions for your schema fields
const resolvers = {
  Query: {
    pub: async (_,{pid}) => {
      //let paper: IPubs[] = pubsC.find({id:pid})
      let paper: IPubs[] = pubsC.findOne({id:pid})
      return paper
    },
    pubs: async () => {
      //let paper: IPubs[] = pubsC.find({id:pid})
      let paper: IPubs[] = pubsC.find()
      return paper
    },
    pquery: async (_,{qry}) => {
      let q: IPubs[] = pubsC.find({'title': {'$regex':qry}});
      return q
    }
  },
};

const server = new ApolloServer({ typeDefs, resolvers });

const app = new Koa();
server.applyMiddleware({ app });

const port = 3000;
const host = 'localhost';

app.listen(port, host, () =>
  console.log(`🚀 Server ready at http://${host}:${port}${server.graphqlPath}`),
);
