var players = {};
var SPRITE_SIZE = 16;

$(document).ready(function() {
    var game = new Phaser.Game(1024, 768, Phaser.AUTO, 'game', { preload: preload, create: create, update: update });

    function preload() {
        game.load.spritesheet('characters', 'sprites.png', SPRITE_SIZE, SPRITE_SIZE);
        game.load.spritesheet('items', 'items.png', SPRITE_SIZE, SPRITE_SIZE);
    }

    function create() {
        var dagger = game.add.sprite(SPRITE_SIZE, SPRITE_SIZE, 'items', 1);

        keyW = game.input.keyboard.addKey(Phaser.Keyboard.UP);
        keyW.onDown.add(movePlayer, this, 0, "n");
        keyA = game.input.keyboard.addKey(Phaser.Keyboard.LEFT);
        keyA.onDown.add(movePlayer, this, 0, "e");
        keyS = game.input.keyboard.addKey(Phaser.Keyboard.DOWN);
        keyS.onDown.add(movePlayer, this, 0, "s");
        keyD = game.input.keyboard.addKey(Phaser.Keyboard.RIGHT);
        keyD.onDown.add(movePlayer, this, 0, "w");

        prepare(game);
    }

    function update() {

    }
});

function movePlayer(ctx, direction) {
    $.post("http://localhost:3000/players/" + localStorage.playerID + "/move/" + direction, function() {});
}

function move(sprite, position) {
    sprite.x = position.x * SPRITE_SIZE;
    sprite.y = position.y * SPRITE_SIZE;
}

function prepare(game) {
    var playerID;
    if (localStorage.playerID) {
        playerID = localStorage.playerID;
        $("#player-id").html(playerID);
    } else {
        $("#player-join").show();
    }

    $.get("http://localhost:3000/players", function(players) {
        for (var i = 0; i < players.length; i++) {
            spawn_player(game, players[i]);
        }
    });

    var source = new EventSource("http://localhost:3000/players/updates");
    source.onmessage = function(event) {
        var e = JSON.parse(event.data);
        console.log(e);
        switch (e.name) {
            case "PlayerJoined":
                spawn_player(game, e.player);
                break;
            case "PlayerMoved":
                move(players[e.player.id], e.player.position);
                break;
        }
    };
}

function spawn_player(game, player) {
    var playerLabel = game.add.text(-15, -15, player.id, {fill: "#fff", font: "12px"});
    var sprite = game.add.sprite(player.position.x * SPRITE_SIZE, player.position.y * SPRITE_SIZE, 'characters', 1);
    sprite.addChild(playerLabel);
    players[player.id] = sprite;
}

function join() {
    var playerID = $("#player-name").val();
    $.post("http://localhost:3000/players/" + playerID, function() {
        localStorage.playerID = playerID;
        $("#player-id").html(playerID);
        $("#player-join").hide();
    })
}
