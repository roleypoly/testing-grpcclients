import React from 'react'
import { DiscordClient, Empty, Guild } from '@roleypoly/rpc/discord'

const discord = new DiscordClient("https://localhost:6777")

interface AppState {
  guilds: Guild[]
}

export default class App extends React.Component<{}, AppState> {
  state = {
    guilds: []
  }

  async componentDidMount () {
    const servers = await discord.listServers(new Empty())
    this.setState({
      guilds: servers.getGuildsList(),
      // guilds: servers.guilds
    })
  }

  render () {
    return <div>
      { this.state.guilds.map((guild: Guild, i) => <p key={i}>{guild.getName()} - {guild.getMembercount()} members</p>) }
    </div>
  }
}
